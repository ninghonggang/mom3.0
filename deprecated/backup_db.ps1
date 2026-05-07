$env:PGPASSWORD = "postgres"
$output = & "C:\Program Files\PostgreSQL\15\bin\pg_dump.exe" -h localhost -U postgres -d mom3 -f "D:\nhgProgram\mom3.0\mom-server\data\mom3_backup.sql"
if (Test-Path "D:\nhgProgram\mom3.0\mom-server\data\mom3_backup.sql") {
    $size = (Get-Item "D:\nhgProgram\mom3.0\mom-server\data\mom3_backup.sql").Length
    Write-Host "Backup created successfully. Size: $size bytes"
} else {
    Write-Host "Backup failed"
    Write-Host $output
}
