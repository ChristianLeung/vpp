*** Variables ***
${KUBE_CLUSTER_1_NODES}            2
${KUBE_CLUSTER_1_VM_1_PUBLIC_IP}   192.168.251.101
${KUBE_CLUSTER_1_VM_1_LOCAL_IP}    192.168.251.101
${KUBE_CLUSTER_1_VM_1_HOST_NAME}   vm1
${KUBE_CLUSTER_1_VM_1_USER}        localadmin
${KUBE_CLUSTER_1_VM_1_PSWD}        cisco123
${KUBE_CLUSTER_1_VM_1_ROLE}        master
${KUBE_CLUSTER_1_VM_1_LABEL}       client_node
${KUBE_CLUSTER_1_VM_2_PUBLIC_IP}   192.168.251.102
${KUBE_CLUSTER_1_VM_2_LOCAL_IP}    192.168.251.102
${KUBE_CLUSTER_1_VM_2_HOST_NAME}   vm2
${KUBE_CLUSTER_1_VM_2_USER}        localadmin
${KUBE_CLUSTER_1_VM_2_PSWD}        cisco123
${KUBE_CLUSTER_1_VM_2_ROLE}        slave
${KUBE_CLUSTER_1_VM_2_LABEL}       server_node
${KUBE_CLUSTER_1_DOCKER_COMMAND}   docker

${KUBE_CLUSTER_2_NODES}            1
${KUBE_CLUSTER_2_VM_1_PUBLIC_IP}   192.168.251.103
${KUBE_CLUSTER_2_VM_1_LOCAL_IP}    192.168.251.103
${KUBE_CLUSTER_2_VM_1_HOST_NAME}   vm3
${KUBE_CLUSTER_2_VM_1_USER}        localadmin
${KUBE_CLUSTER_2_VM_1_PSWD}        cisco123
${KUBE_CLUSTER_2_DOCKER_COMMAND}   docker

${KUBE_CLUSTER_3_NODES}            1
${KUBE_CLUSTER_3_VM_1_PUBLIC_IP}   192.168.251.104
${KUBE_CLUSTER_3_VM_1_LOCAL_IP}    192.168.251.104
${KUBE_CLUSTER_3_VM_1_HOST_NAME}   vm4
${KUBE_CLUSTER_3_VM_1_USER}        localadmin
${KUBE_CLUSTER_3_VM_1_PSWD}        cisco123
${KUBE_CLUSTER_3_DOCKER_COMMAND}   docker

${KUBE_CLUSTER_4_NODES}            1
${KUBE_CLUSTER_4_VM_1_PUBLIC_IP}   192.168.252.101
${KUBE_CLUSTER_4_VM_1_LOCAL_IP}    192.168.252.101
${KUBE_CLUSTER_4_VM_1_USER}        localadmin
${KUBE_CLUSTER_4_VM_1_PSWD}        cisco123
${KUBE_CLUSTER_4_DOCKER_COMMAND}   docker

${KUBE_CLUSTER_5_NODES}            2
${KUBE_CLUSTER_5_VM_1_PUBLIC_IP}   192.168.251.105
${KUBE_CLUSTER_5_VM_1_LOCAL_IP}    192.168.251.105
${KUBE_CLUSTER_5_VM_1_HOST_NAME}   vm5
${KUBE_CLUSTER_5_VM_1_USER}        localadmin
${KUBE_CLUSTER_5_VM_1_PSWD}        cisco123
${KUBE_CLUSTER_5_VM_1_ROLE}        master
${KUBE_CLUSTER_5_VM_1_LABEL}       client_node
${KUBE_CLUSTER_5_VM_2_PUBLIC_IP}   192.168.251.106
${KUBE_CLUSTER_5_VM_2_LOCAL_IP}    192.168.251.106
${KUBE_CLUSTER_5_VM_2_HOST_NAME}   vm6
${KUBE_CLUSTER_5_VM_2_USER}        localadmin
${KUBE_CLUSTER_5_VM_2_PSWD}        cisco123
${KUBE_CLUSTER_5_VM_2_ROLE}        slave
${KUBE_CLUSTER_5_VM_2_LABEL}       server_node
${KUBE_CLUSTER_5_DOCKER_COMMAND}   docker

# Other variables
${RESULTS_FOLDER}                  results
${TEST_DATA_FOLDER}                test_data
${SSH_READ_DELAY}                  3
