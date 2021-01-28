

call gcloud functions deploy CreatePostFunction --entry-point CreatePostFunction --runtime go113 --max-instances 2 --trigger-http --allow-unauthenticated
call gcloud functions deploy ToggleLikePostFunction --entry-point ToggleLikePostFunction --runtime go113 --max-instances 2 --trigger-http --allow-unauthenticated
call gcloud functions deploy ToggleRetweetPostFunction --entry-point ToggleRetweetPostFunction --runtime go113 --max-instances 2 --trigger-http --allow-unauthenticated
call gcloud functions deploy GetPostByIdFunction --entry-point GetPostByIdFunction --runtime go113 --max-instances 2 --trigger-http --allow-unauthenticated

