export JAVA_HOME=$(/usr/libexec/java_home)

export VFLOWPATH="$GOPATH/src/github.wdf.sap.corp/velocity/vflow"

export GONOPROXY="*.sap.corp"
export GOPRIVATE="*.sap.corp"

export CONN_MANAGEMENT_ENDPOINT=http://localhost:3000

export PYTHONPATH=$VFLOWPATH/src/stdlib/subdevkits/pythree_subengine/:$VFLOWPATH/src/stdlib/subengines/com/sap/

export VFLOW_IM_TOKEN="nananana"

export SAP_DI_VRELEASE_DELIVERY_DI_VFLOW_RESILIENT_PIPELINES=true

export HANA_HOST="localhost"
export HANA_PORT="30017"
export HANA_VFLOW_USER="SYSTEM"
export HANA_VFLOW_USER_PASSWORD=""
export HANA_SERVER="FALSE"
export KUBECONFIG="$HOME/KUBECONFIG.yaml"

export GCLOUD_KEY_FILE="$HOME/key-gcloud.json"
export DI_USER_TENANT_PASSWORD=''
export DI_USER_TENANT_NAME="default"
export DI_USER_TENANT_USER="system"
export DI_SYSTEM_TENANT_PASSWORD=''
export DI_SYSTEM_TENANT_NAME="system"
export DI_SYSTEM_TENANT_USER="system"
export DI_LAUNCHPAD_ENDPOINT=""
export DOCKER_TAG=""
export DOCKER_REGISTRY=""

export USER_TENANT_PASSWORD=$DI_USER_TENANT_PASSWORD
export USER_TENANT_NAME=$DI_USER_TENANT_NAME
export USER_TENANT_USER=$DI_USER_TENANT_USER
export SYSTEM_TENANT_PASSWORD=$DI_SYSTEM_TENANT_PASSWORD
export SYSTEM_TENANT_NAME=$DI_SYSTEM_TENANT_NAME
export SYSTEM_TENANT_USER=$DI_SYSTEM_TENANT_USER
export LAUNCHPAD_ENDPOINT=$DI_LAUNCHPAD_ENDPOINT

alias kegrafset="kubectl edit configmap diagnostics-grafana-settings"

vflow-build () 
{ 
    ( cd $VFLOWPATH/src && make vflow "$@" )
}

vflow-run() {
    vflow-build && vflow-run
}

vflow-nobuild-run () 
{ 
    ( cd $VFLOWPATH/src/stdlib && vflow --local --externalui=ui "$@" )
}


vflow-lint ()
{
    golangci-lint run -c $VFLOWPATH/.golangci.yml --deadline=60m --max-issues-per-linter=0 --max-same-issues=0
}

vctl-login-system() 
{
    vctl login -i "$LAUNCHPAD_ENDPOINT" "$SYSTEM_TENANT_USER" "$SYSTEM_TENANT_NAME" -p "$SYSTEM_TENANT_PASSWORD"
}

vctl-login-user() 
{
    vctl login -i "$LAUNCHPAD_ENDPOINT" "$USER_TENANT_USER" "$USER_TENANT_NAME" -p "$USER_TENANT_PASSWORD"
}

vctl-restart-vflow() {
    vctl-login-system && vctl scheduler restart vflow
}

vctl-get-version() {
    vctl scheduler get-template vflow | grep vflow | awk -F " " '{print $3}'
}

patch-v-cluster-app() {
    vctl-login-system
    DOCKER_TAG="$(vctl-get-version)"
    start-docker && cd $VFLOWPATH/src && make docker_image_vflow-cluster-app && make docker_push_vflow-cluster-app && vctl-restart-vflow
}

# The next line updates PATH for the Google Cloud SDK.
if [ -f "$HOME/google-cloud-sdk/path.zsh.inc" ]; then . "$HOME/google-cloud-sdk/path.zsh.inc"; fi

# The next line enables shell command completion for gcloud.
if [ -f "$HOME/google-cloud-sdk/completion.zsh.inc" ]; then . "$HOME/google-cloud-sdk/completion.zsh.inc"; fi

# >>> xmake >>>
[[ -s "$HOME/.xmake/profile" ]] && source "$HOME/.xmake/profile" # load xmake profile
# <<< xmake <<<