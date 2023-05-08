package app

import (
	"policy/docs"
	controller_admin "policy/modules/agent/controllers"
	controller_location "policy/modules/locations/controllers"
	controller_policy "policy/modules/policies/controllers"
	controller_users "policy/modules/users/controllers"
	controller_offers "policy/modules/offers/controllers"
	cloud_bucket "policy/modules/cloud_bucket"




)

func mapUrls() {
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group(docs.SwaggerInfo.BasePath)
	{
		agentAPi := v1.Group("/agent")
		agentAPi.POST("", controller_admin.AddUserToAdmin)
		agentAPi.GET("", controller_admin.GetAllUsersinAdmin)
		agentAPi.GET(":user_id", controller_admin.GetUsersById)

		locationsAPI := v1.Group("/locations")
		locationsAPI.POST("", controller_location.AddLocation)
		locationsAPI.GET("", controller_location.GetAllLocations)
		locationsAPI.GET(":location_id", controller_location.GetLocationById)

		policiesAPI := v1.Group("/policies")
		policiesAPI.POST("", controller_policy.AddPolicy)
		policiesAPI.GET("", controller_policy.GetAllPolicies)
		policiesAPI.GET(":policy_id", controller_policy.GetPolicyById)

		cloudStorageAPI := v1.Group("/file_upload")
		cloudStorageAPI.POST("", cloud_bucket.HandleFileUploadToBucket)

		usersApi := v1.Group("/users")
		usersApi.POST("", controller_users.AddUser)
		usersApi.GET("", controller_users.GetAllUsers)
		usersApi.GET(":user_id", controller_users.GetUserById)

		offersApi := v1.Group("/offers")
		offersApi.POST("", controller_offers.AddOffer)
		offersApi.GET("", controller_offers.GetAllOffers)
		offersApi.GET(":offer_id", controller_offers.GetOffersById)
	}
}
