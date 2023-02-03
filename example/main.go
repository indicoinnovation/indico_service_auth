package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	indicoserviceauth "github.com/INDICO-INNOVATION/indico_service_auth"
	serviceaccounts "github.com/INDICO-INNOVATION/indico_service_auth/client/service_account"
)

func main() {
	client, ctx, err := indicoserviceauth.NewClient()
	if err != nil {
		log.Fatalf(err.Error())
	}

	testBackoffice(client, ctx)

	// generateAndValidate(client, ctx)
	// validateThird(client, ctx, "643863")
}

// func generateAndValidate(client *indicoserviceauth.Client, ctx context.Context) {
// 	response, err := client.GenerateOTP(ctx)
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	fmt.Println("Generate OTP Response:")
// 	fmt.Printf("%+v\n\n", response)

// 	responsev, err := client.ValidateOTP(ctx, response.Token, true)
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	fmt.Println("Validate OTP Response:")
// 	fmt.Printf("%+v\n\n", responsev)
// }

// func validateThird(client *indicoserviceauth.Client, ctx context.Context, token string) {
// 	responsev, err := client.ValidateOTP(ctx, token, false)
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	fmt.Println("Validate Third Party OTP Response:")
// 	fmt.Printf("%+v\n\n", responsev)

// 	ctx.Done()
// }

func testBackoffice(client *indicoserviceauth.Client, ctx context.Context) {
	testServiceAccounts(client, ctx)

	// response, err := client.ListResources(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(response)

	// fmt.Println("")

	// response2, err := client.CreateClient(ctx, "Anthony", "user")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(response2)

	// fmt.Println("")

	// response3, err := client.CreateResource(ctx, "use", "Use stuff")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(response3)

	// fmt.Println("")

	// testCreateResourceScope(client, ctx)

	// response5, err := client.CreateServiceAccount(ctx, "Tony", "tony", "he likes food")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(response5)

	// fmt.Println("")

	// response6, err := client.GetResourceScope(ctx, 5)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(response6)
	// }

	// func testCreateResourceScope(client *indicoserviceauth.Client, ctx context.Context) {
	// 	payload := &indicoserviceauth.ResourceScopeList{
	// 		Data: []*indicoserviceauth.ResourceScope{
	// 			{
	// 				Name:        "role_1",
	// 				Label:       "Role_1",
	// 				Description: "Access to role 1",
	// 				ResourceID:  1,
	// 			},
	// 			{
	// 				Name:        "role_2",
	// 				Label:       "Role_2",
	// 				Description: "Access to role 2",
	// 				ResourceID:  1,
	// 			},
	// 		},
	// 	}

	// response, err := client.CreateResourceScope(ctx, payload)
	//
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	// fmt.Println(response)
	// fmt.Println("")
}

func testServiceAccounts(client *indicoserviceauth.Client, ctx context.Context) {
	var serviceAccounts []*serviceaccounts.ServiceAccount
	var serviceAccountsToDelete []string

	for _, i := range []string{"1", "2", "3"} {
		serviceAccount, err := client.CreateServiceAccount(ctx, "Teste"+i, "teste"+i, "sa creation lib test")
		if err != nil {
			log.Fatalf(err.Error())
		}
		serviceAccounts = append(serviceAccounts, serviceAccount)
		serviceAccountsToDelete = append(serviceAccountsToDelete, strconv.Itoa(int(serviceAccount.ServiceAccountId)))
	}

	fmt.Println("Created SAs:")
	fmt.Println(serviceAccounts)
	fmt.Println("")

	serviceAccounts, err := client.DeleteServiceAccount(ctx, serviceAccountsToDelete)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Deleted SA:")
	fmt.Printf("%+v", serviceAccounts)
	fmt.Println("")

	serviceAccounts, err = client.ListServiceAccounts(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("SA Listing:")
	fmt.Printf("%+v \n", serviceAccounts)
	fmt.Println("")

	credentials, err := client.GenerateCredentials(ctx, strconv.Itoa(int(serviceAccounts[0].ServiceAccountId)))
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Credentials Generated:")
	fmt.Printf("%+v \n", credentials)
	fmt.Println("")
}
