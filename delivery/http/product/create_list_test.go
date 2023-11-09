package product

// TODO: do test it later because need create csv file function
// func TestCreateList(t *testing.T) {
// 	assertion := assert.New(t)
// 	db := database.InitDatabase()
// 	defer db.TruncateTables()

// 	repo := repository.New(db.GetClient)
// 	r := Route{
// 		UseCase: usecase.New(repo, nil),
// 	}

// 	accountID, producerID, err := SetUpForeignKeyData(db)
// 	assertion.NoError(err)

// 	userPrinciple := &token.JWTClaimCustom{
// 		SessionID: uuid.New(),
// 		User: token.UserInfo{
// 			Username: gofakeit.Name(),
// 			ID:       accountID,
// 			Email:    gofakeit.Email(),
// 		},
// 	}
// 	monkey.Patch(contexts.GetUserPrincipleByContext, func(context.Context) *token.JWTClaimCustom {
// 		return userPrinciple
// 	})

// 	defer monkey.UnpatchAll()

// 	// good case
// 	{
// 		// Arrange

// 		// Act

// 		// Assert
// 	}

// 	// bad case
// 	{
// 		// Arrange

// 		// Act

// 		// Assert
// 	}
// }
