

Installalas: 
docker-compose up

TODO: ha a frontend data.services.ts-ben beallitom a kontener hostname-jet a backendnek (shorter) akkor valamiert nem tud csatlakozni a frontend a backendhez. Ha ipt allitok be akkor jo...
                (private API_SERVER = "http://172.22.0.4:3000";  -  private API_SERVER = "http://shorter:3000";)   
Ez miatt lehet nalad nem fog mukodni, kiveve ha beallitod a backend ip-jet abban a fileban... valamiert nem oldja fel a "shorter"-t
ranezek majd kesobb, hogy ez miert lehet


---------------------------------------------------------------------------------------------------------------------------
Dokumentacio: 

A directory strukturaja:
web - angular frontend 
docker - Dockerfileok a kontener kesziteshez
openapi - a feladathoz szukseges api hivasok  generalasa
shorter - backend 
sql - a psql kontener felallasa utan vegrehajtott sql utasitasok


BACKEND: 
A feladatot hexagonal arhitekturaban oldottam meg.
Ez az arhitektura kicsit overkill ehez a feladathoz, de nekem tetszik. (atlathatosag, ujrafelhasznalas, konnyebb tesztelhetoseg)

A driver (adapter) itt az internal/httphandler
driven (adapter) internal/repositories

internal/core/services-ben vannak a query, command, event interakciok megvalositasa 
ezek felelosek a domain logic es a repo kezeleseert

A pkg dir-ben vannak a konnyen mashol is hasznalhato csomagok. (log,error,httpserver kiegeszitok)

internal/httphandler - felelos a request ellenorzesert, response gyartasert

------------------------------------------------------------------------------------------
FRONTEND:
about,home,post component
home - megnyitasa betolti az osszes url-hash part
post - itt lehet addolni 
backend api hivasokat az angluar httpclient vegzi a data services-ben




shorter: 
+---shorter
|   |   go.mod
|   |   go.sum
|   |   
|   +---internal
|   |   +---handlers
|   |   |   \---httphandler
|   |   |           httphandler.go   
|   |   |           httphandler_gen.go
|   |   |           
|   |   +---core
|   |   |   +---services
|   |   |   |   +---event
|   |   |   |   +---command
|   |   |   |   |       commandmakeurlhash.go
|   |   |   |   |       commanddeleteurl.go
|   |   |   |   |       
|   |   |   |   \---query
|   |   |   |           querygeturl.go
|   |   |   |           queryurlswidthhash.go
|   |   |   |           
|   |   |   +---domain
|   |   |   |       domain.go
|   |   |   |       domain_gen.go
|   |   |   |       
|   |   |   \---ports
|   |   |           portrepo.go
|   |   |           portservice.go
|   |   |           
|   |   \---repositories
|   |       +---psqlrepo
|   |       |       psqlrepo.go
|   |       |       
|   |       \---memrepo
|   |               memrepo.go
|   |               
|   +---cmd
|   |   |   go.mod
|   |   |   
|   |   \---shorter
|   |           shorter.go
|   |           cmd
|   |           go.mod
|   |           go.sum
|   |           
|   \---pkg
|       +---httpserver
|       |       httpserver.go
|       |       
|       +---errors
|       |   +---httperror
|       |   |       httperror.go
|       |   |       
|       |   \---exterror
|       |           exterror.go
|       |           
|       \---logs
|           +---extlog
|           |       extlog.go
|           |       
|           \---httplog
|                   echohttplog.go
|                   
\---openapi
        shorter_gen.go
        shorter.yaml
        shorter_dom_gen.go
