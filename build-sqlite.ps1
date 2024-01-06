#!/usr/bin/env pwsh

$cwd = Get-Location

Set-Location $PSScriptRoot

go build -tags sqlite

Set-Location $cwd
