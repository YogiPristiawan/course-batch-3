export $(grep -v '^#' .env | sed -r '/^\s*$/d' | xargs)