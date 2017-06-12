@ECHO OFF
SETLOCAL

REM Directories
SET SRCDIR=%CD%

REM @SET ALLDIRS!=	find ${SRCDIR} -type d -not -path "./.hg*"

REM Packages
SET ROOTPKG=github.com/veandco/go-sdl2
SET PACKAGES=sdl img mix ttf

SET GO=go

@IF "%~1" == "" GOTO :all
@GOTO :%~1

:all
CALL :packages
GOTO :eof

:test
CALL :packages
FOR %%p IN (%PACKAGES%) DO (
	ECHO Testing %ROOTPKG%/%%p...
	%GO% test -v %ROOTPKG%/%%p
)
GOTO :eof

:bench
CALL :packages
FOR %%p IN (%PACKAGES%) DO (
	ECHO Benchmarking %ROOTPKG%/%%p...
	%GO% test -bench=. -cpu 1,2,3,4 %ROOTPKG%/%%p
)
GOTO :eof

:packages
FOR %%p IN (%PACKAGES%) DO (
	ECHO Building package %ROOTPKG%/%%p...
	%GO% build %ROOTPKG%/%%p
)
GOTO :eof

ENDLOCAL
ECHO ON
