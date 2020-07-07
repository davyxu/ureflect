# ureflect
golang reflect used like c#



```go
	prof := &Profile{
		Name: "cat",
		Age:  18,
	}

	tProfile := TypeOf(prof)	

	fieldAge := tProfile.FieldByName("Age")	
	
	fmt.Println(fieldAge.Int32(prof))

    fmt.Println(tProfile.FieldByName("Name").String(PointerOf(prof)))	

	tProfile.Field(0).SetString(prof, "dog")
	

	tProfile.Field(1).SetInt32(prof, 1)
	
	
```