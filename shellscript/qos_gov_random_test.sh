#!/bin/bash

testTimes=500

HOSTS=(192.168.1.200 192.168.1.200 192.168.1.200 192.168.1.200)
TYPE=(submit_proposal submit_proposal vote deposit vote submit_proposal vote vote submit_proposal deposit vote submit_proposal deposit deposit)
#TYPE=(submit_proposal submit_proposal submit_proposal)
PROPOSAL_TYPE=(Text ParameterChange TaxUsage)
OPTIONS=(yes no abstain nowithveto yes yes yes no no abstain)

QOSCLI="~/gocode/bin/qoscli"
QOSPWD="11111111"

SSHCOMMAND="sshpass -pvagrant ssh vagrant@HOST \"COMMAND\""


#run_command host command
#ex. run_command 192.168.1.224 "$QOSCLI q validators --indent"
run_command(){
  c=${SSHCOMMAND/HOST/$1}
  c=${c/COMMAND/$2}
  eval $c
}

random(){
  ((r=$RANDOM%$1))
  echo $r
}

VALIDATORS=()
DELEGATORS=()
init_validators(){
  index=0
  all_validator=$(run_command ${HOSTS[0]} "$QOSCLI q validators --indent" | grep 'owner'|awk '{print $2}'|grep -Po '[a-z0-9A-Z]*')
  for i in $(echo $all_validator | tr ' ' '\n')
  do
    VALIDATORS[$index]=$i
    DELEGATORS[$index]=$i
    ((index=index+1))
  done
}

GUARDIANS=()
init_guardians(){
  indexg=0
  all_guardian=$(run_command ${HOSTS[0]} "$QOSCLI query guardians --indent" | grep 'address' | awk '{print $2}'| grep -Po '[a-z0-9A-Z]*')
  for j in $(echo $all_guardian | tr ' ' '\n')
  do
    GUARDIANS[$indexg]=$j
    ((indexg=indexg+1))
  done
}

#1. init all operator users
init_validators

init_guardians

#$host $proposer $deposit 
submit_proposal_text(){
  run_command $1 "echo $QOSPWD | $QOSCLI tx submit-proposal --title 'myproposal' --proposal-type 'Text' --proposer $2 --deposit $3 --description 'the first proposal'"
}
#$host $proposer $deposit 
submit_proposal_parameterchange(){
  run_command $1 "echo $QOSPWD | $QOSCLI tx submit-proposal --title 'myproposal' --proposal-type 'ParameterChange' --proposer $2 --deposit $3 --description 'the first proposal' --params gov:min_deposit:1000"
}
#$host $proposer $deposit $dest_address
submit_proposal_taxusage(){
  run_command $1 "echo $QOSPWD | $QOSCLI tx submit-proposal --title 'myproposal' --proposal-type 'TaxUsage' --proposer $2 --deposit $3 --description 'the first proposal' --dest-address $4 --percent 0.8"
}


#$host $proposal_id $depositor $amount
deposit(){
  run_command $1 "echo $QOSPWD | $QOSCLI tx deposit --proposal-id $2 --depositor $3 --amount $4"
}

#$host $proposal_id $voter $options
vote(){
  run_command $1 "echo $QOSPWD | $QOSCLI tx vote --proposal-id $2 --voter $3 --option $4"
}


#2. main
count=0
while [[ $count -lt $testTimes ]]; do

  type=${TYPE[$(random ${#TYPE[@]})]}
  host=${HOSTS[$(random ${#HOSTS[@]})]}
  ((deposit=$RANDOM%100+1))
  proposer=${VALIDATORS[$(random ${#VALIDATORS[@]})]}
  depositor=${DELEGATORS[$(random ${#DELEGATORS[@]})]}
  proposal_type=${PROPOSAL_TYPE[$(random ${#PROPOSAL_TYPE[@]})]}
  guardian=${GUARDIANS[$(random ${#GUARDIANS[@]})]}
  ((proposal_id=$RANDOM%10+1))
  options=${OPTIONS[$(random ${#OPTIONS[@]})]}

  case $type in
    "submit_proposal")
      case $proposal_type in
      "Text")
        submit_proposal_text $host $proposer $deposit 
      ;;
      "ParameterChange")
        submit_proposal_parameterchange $host $proposer $deposit 
      ;;
      "TaxUsage")
        submit_proposal_taxusage $host $proposer $deposit $guardian
      ;;
      esac
      
      ;;
    "deposit")
      deposit $host $proposal_id $depositor $deposit
      ;;
    "vote")
      vote $host $proposal_id $depositor $options
      ;;
  esac


  ((count=count+1))
  # ((tt=$RANDOM%10+1))
  # sleep $tt

done
