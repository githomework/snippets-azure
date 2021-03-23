
**Change yml on Azure**
az webapp config container set --resource-group <rg-aea-dev-xxx>   --name <name> --multicontainer-config-type compose --multicontainer-config-file compose-wordpress.yml
