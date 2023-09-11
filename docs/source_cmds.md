# Source Commands

`hava source -h`

```bash
Create/List/Delete/Update sources to Hava for different providers (AWS/Azure/Google Cloud)

Usage:
  hava source [flags]
  hava source [command]

Available Commands:
  create      Create sources to Hava
  delete      Delete sources to Hava
  list        List sources of Hava
  sync        Sync sources to Hava
  update      Update sources to Hava

Flags:
      --all               hava source --all
  -h, --help              help for source
      --sourceId string   sourceId of AWS|Azure|GCP source
```

## Source List all

`hava source list`

```
╭───┬─────────────┬──────────────────────────────────────┬──────────────────────┬────────────┬────────┬────────────────────╮
│   │ DISPLAYNAME │ ID                                   │ INFO                 │ NAME       │ STATE  │ TYPE               │
├───┼─────────────┼──────────────────────────────────────┼──────────────────────┼────────────┼────────┼────────────────────┤
│ 1 │ devTestAWS  │ 8eb192e2-9beb-466b-ae14-c05fc8403cf4 │ AKIAIOSFODNN7EXAMPLE │ devTestAWS │ active │ Sources::AWS::Keys │
╰───┴─────────────┴──────────────────────────────────────┴──────────────────────┴────────────┴────────┴────────────────────╯
```

## Source List SourceID

`hava source list --source-id a58b7cb1-f9da-42ad-9fc1-8dc61b0d3e38`


## Source Sync SourceID

`hava source sync a58b7cb1-f9da-42ad-9fc1-8dc61b0d3e38`


## Source Create AWS (Using Access Keys)

`hava source create aws --name dev --access-key $AWS_ACCESS_KEY_ID --secret-key $AWS_SECRET_ACCESS_KEY`

```bash
[INFO]   Created AWS Source for the following source:

╭───┬─────────────┬──────────────────────────────────────┬──────────────────────┬────────────┬────────┬────────────────────╮
│   │ DISPLAYNAME │ ID                                   │ INFO                 │ NAME       │ STATE  │ TYPE               │
├───┼─────────────┼──────────────────────────────────────┼──────────────────────┼────────────┼────────┼────────────────────┤
│ 1 │ devTestAWS  │ 040d5b5e-03b9-4343-9393-8ad0794512f4 │ AKIAIOSFODNN7EXAMPLE │ devTestAWS │ queued │ Sources::AWS::Keys │
╰───┴─────────────┴──────────────────────────────────────┴──────────────────────┴────────────┴────────┴────────────────────╯
```

## Source Create AWS (Using Cross Account Role)

`hava source create aws --name devCAR --role-arn $AWS_CROSS_ACCOUNT_ROLE`

```bash
[INFO]   Created AWS Source for the following source:

╭───┬─────────────┬──────────────────────────────────────┬────────────────────────────────────────┬────────┬────────┬────────────────────────────────╮
│   │ DISPLAYNAME │ ID                                   │ INFO                                   │ NAME   │ STATE  │ TYPE                           │
├───┼─────────────┼──────────────────────────────────────┼────────────────────────────────────────┼────────┼────────┼────────────────────────────────┤
│ 1 │ devCAR      │ 31f9b6a6-4e48-400a-8ced-5bfc2a1aacc2 │ arn:aws:iam::123456789012:role/HavaCAR │ devCAR │ queued │ Sources::AWS::CrossAccountRole │
╰───┴─────────────┴──────────────────────────────────────┴────────────────────────────────────────┴────────┴────────┴────────────────────────────────╯
```

## Source Create GCP

`hava source create gcp  --name GCPDev --configFile $GCP_ENCODED_FILE`

```bash
[INFO]   Created GCP Source for the following source:

╭───┬─────────────┬──────────────────────────────────────┬──────────────────┬────────┬────────┬─────────────────────────────────────────╮
│   │ DISPLAYNAME │ ID                                   │ INFO             │ NAME   │ STATE  │ TYPE                                    │
├───┼─────────────┼──────────────────────────────────────┼──────────────────┼────────┼────────┼─────────────────────────────────────────┤
│ 1 │ GCPDev      │ 8818d864-46e9-4b8d-acdd-a31d8936e62f │ credentials.json │ GCPDev │ queued │ Sources::GCP::ServiceAccountCredentials │
╰───┴─────────────┴──────────────────────────────────────┴──────────────────┴────────┴────────┴─────────────────────────────────────────╯
```

## Source Create Azure

`hava create source azure --name AzureDev --client-id $ARM_CLIENT_ID --tenant-id $ARM_TENANT_ID --subscription-id $ARM_SUBSCRIPTION_ID`


## Source Delete

`hava source delete 22872411-20e8-4b6e-aa46-41866c9c1897`


## Source Update AWS|AZURE|GCP

`hava source update gcp --name GCPDevChange --source-id f2a26440-10bf-43d1-9742-8361de30590f`

