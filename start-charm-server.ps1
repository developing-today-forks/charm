#!/usr/bin/env pwsh

$cwd = Get-Location

."$PSScriptRoot/build-libsql.ps1"

$env:CHARM_SERVER_DB_DRIVER="libsql"

if ([string]::IsNullOrEmpty($env:TURSO_HOST)) {
    throw "TURSO_HOST environment variable must be set"
}
if ([string]::IsNullOrEmpty($env:TURSO_AUTH_TOKEN)) {
   throw "TURSO_AUTH_TOKEN environment variable must be set"
}
$env:CHARM_SERVER_DB_DATA_SOURCE="libsql://${env:TURSO_HOST}?authToken=${env:TURSO_AUTH_TOKEN}"

Set-Location $PSScriptRoot

charm serve

Set-Location $cwd
