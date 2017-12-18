USE [master]
RESTORE DATABASE [db-passport-2.5.191] FROM  DISK = N'C:\Program Files\Microsoft SQL Server\MSSQL11.SQL2012R1DEV\MSSQL\Backup\db-passport-2.5.187-baseline.bak' WITH  FILE = 1,  MOVE N'matter_management' TO N'C:\Program Files\Microsoft SQL Server\MSSQL11.SQL2012R1DEV\MSSQL\DATA\db-passport-2.5.191.MDF',  MOVE N'matter_management_log' TO N'C:\Program Files\Microsoft SQL Server\MSSQL11.SQL2012R1DEV\MSSQL\DATA\db-passport-2.5.191_1.LDF',  NOUNLOAD,  REPLACE,  STATS = 5

GO

