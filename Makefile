ts_account_service:
  hz update -idl idl/$(service).proto --customize_package=./template/package.yaml

all:
	make ts_account_service