

call gcloud functions deploy CreateUser --entry-point CreateUser --runtime go113 --max-instances 2 --trigger-http --allow-unauthenticated
call gcloud functions deploy UpdateUserTagFunction --entry-point UpdateUserTagFunction --runtime go113 --max-instances 2 --trigger-http --allow-unauthenticated
call gcloud functions deploy CheckDoc --entry-point CheckDoc --runtime go113 --max-instances 2 --trigger-http --allow-unauthenticated
call gcloud functions deploy ToggleFollow --entry-point ToggleFollow --runtime go113 --max-instances 2 --trigger-http --allow-unauthenticated
call gcloud functions deploy GetUserFunction --entry-point GetUserFunction --runtime go113 --max-instances 2 --trigger-http --allow-unauthenticated

