# This is an http test server to test API calls

The Payload struct is
```
{
"bundle":"teams_trial",
"UUID": "pooltenant27111263", // Update the datetime stamp and then keep incrementing by 1
"name": "pooltenant27111263, // Update the datetime stamp and then keep incrementing by 1
"endDate": null,
"options": {
"professionals": "1",
"solvers": "25",
"additionalProfessionals": 0,
"additionalSolvers": 0
  }
}
```
and it is defined in the server as 
```
type Payload struct {
	Bundle  string  `json:"bundle"`
	UUID    string  `json:"UUID"`
	Name    string  `json:"name"`
	EndDate *string `json:"endDate"`
	Options Options `json:"options"`
}
```