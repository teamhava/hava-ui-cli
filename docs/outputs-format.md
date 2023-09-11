## Output Formats



### Table Output
A table is the default format when outputting information about hava sources

```bash
hava source list                                                                                     
╭───┬───────────────┬──────────────────────────────────────┬──────────────────────┬───────────────┬────────┬─────────────────────────────────────────╮ 
│   │ DISPLAYNAME   │ ID                                   │ INFO                 │ NAME          │ STATE  │ TYPE                                    │
├───┼───────────────┼──────────────────────────────────────┼──────────────────────┼───────────────┼────────┼─────────────────────────────────────────┤
│ 1 │ dev           │ 4f14c115-3b0d-40ea-b075-6df9b2fb81c9 │ AKIAIOSFODNN7EXAMPLE │ dev           │ active │ Sources::AWS::Keys                      │
│ 2 │ GCPDevChange3 │ f2a26440-10bf-43d1-9742-8361de30590f │ credentials.json     │ GCPDevChange3 │ active │ Sources::GCP::ServiceAccountCredentials │
╰───┴───────────────┴──────────────────────────────────────┴──────────────────────┴───────────────┴────────┴─────────────────────────────────────────╯
```

### Json Output


```bash
hava source list --json | jq 
[
  {
    "DisplayName": "dev",
    "Id": "4f14c115-3b0d-40ea-b075-6df9b2fb81c9",
    "Info": "AKIAIOSFODNN7EXAMPLE",
    "Name": "dev",
    "State": "active",
    "Type": "Sources::AWS::Keys"
  },
  {
    "DisplayName": "GCPDevChange3",
    "Id": "f2a26440-10bf-43d1-9742-8361de30590f",
    "Info": "credentials.json",
    "Name": "GCPDevChange3",
    "State": "active",
    "Type": "Sources::GCP::ServiceAccountCredentials"
  }
]
```

### CSV Output

```bash
hava source list --csv       
,DisplayName,Id,Info,Name,State,Type
1,dev,4f14c115-3b0d-40ea-b075-6df9b2fb81c9,AKIAIOSFODNN7EXAMPLE,dev,active,Sources::AWS::Keys
2,GCPDevChange3,f2a26440-10bf-43d1-9742-8361de30590f,credentials.json,GCPDevChange3,active,Sources::GCP::ServiceAccountCredentials
```


### Markdown Output

```bash
hava source list --markdown
| | DisplayName | Id | Info | Name | State | Type |
| ---:| --- | --- | --- | --- | --- | --- |
| 1 | dev | 4f14c115-3b0d-40ea-b075-6df9b2fb81c9 | AKIAIOSFODNN7EXAMPLE | dev | active | Sources::AWS::Keys |
| 2 | GCPDevChange3 | f2a26440-10bf-43d1-9742-8361de30590f | credentials.json | GCPDevChange3 | active | Sources::GCP::ServiceAccountCredentials |
```


### HTML Output

```
hava source list --html
<table class="go-pretty-table">
  <thead>
  <tr>
    <th>&nbsp;</th>
    <th>DisplayName</th>
    <th>Id</th>
    <th>Info</th>
    <th>Name</th>
    <th>State</th>
    <th>Type</th>
  </tr>
  </thead>
  <tbody>
  <tr>
    <td align="right">1</td>
    <td>appGCP</td>
    <td>09409035-87d4-478b-bec0-25349825d390</td>
    <td>credentials.json</td>
    <td>appGCP</td>
    <td>active</td>
    <td>Sources::GCP::ServiceAccountCredentials</td>
  </tr>
  <tr>
    <td align="right">2</td>
    <td>devTestAWS</td>
    <td>a2d8d0c8-0fd4-4483-ab2a-71e4ef58e093</td>
    <td>AKIAIOSFODNN7EXAMPLE</td>
    <td>devTestAWS</td>
    <td>active</td>
    <td>Sources::AWS::Keys</td>
  </tr>
  <tr>
    <td align="right">3</td>
    <td>devTestAWScar</td>
    <td>c3b9d2d7-0cb1-4fd1-92d7-beed7ceb7b22</td>
    <td>arn:aws:iam::777777777777:role/HavaCAR</td>
    <td>devTestAWScar</td>
    <td>active</td>
    <td>Sources::AWS::CrossAccountRole</td>
  </tr>
  <tr>
    <td align="right">4</td>
    <td>devTestAzure</td>
    <td>76903928-8ae0-4ccb-8dee-ede61db9dc5c</td>
    <td>uPfVzmPStDe9ya</td>
    <td>devTestAzure</td>
    <td>inactive</td>
    <td>Sources::Azure::Credentials</td>
  </tr>
  </tbody>
</table>
```
