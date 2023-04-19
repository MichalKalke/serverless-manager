#!/bin/bash

function get_kyma_status () {
	local number=1
	while [[ $number -le 100 ]] ; do
		echo ">--> checking serverless status #$number"
		local STATUS=$(kubectl get serverless -n kyma-system serverless-k3d -o jsonpath='{.status.state}')
		echo "serverless status: ${STATUS:='UNKNOWN'}"
		[[ "$STATUS" == "Ready" ]] && return 0
		sleep 5
        	((number = number + 1))
	done

	kubectl get all --all-namespaces
	exit 1
}

get_kyma_status