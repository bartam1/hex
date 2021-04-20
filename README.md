

Installalas: 
docker-compose up

Ez a hiba nem lepett fel kesobbiekben: (
TODO: ha a frontend data.services.ts-ben beallitom a backend kontener hostname-jet  (shorter), akkor valamiert nem tud csatlakozni a frontend a backendhez. Ha ipt allitok be akkor jo...
                (private API_SERVER = "http://172.22.0.4:3000";  -  private API_SERVER = "http://shorter:3000";)   

Az API_SERVER cim a ".env-frontend" fileban van definialva
A docker-compose-al lehetne allitani fix ip-ket a kontenereknek es akkor az megoldas.
)
---------------------------------------------------------------------------------------------------------------------------
Dokumentacio: 

A directory strukturaja:
web - angular frontend 
docker - Dockerfileok a kontener kesziteshez
openapi - a feladathoz szukseges api hivasok  generalasa
shorter - backend 
sql - a psql kontener felallasa utan vegrehajtott sql utasitasok

A feladattol eltertem a routingnal a /url/{hash} /url/del/{hash} lett a DELETE operacio
Ez az openapi miatt volt szukseges. Termeszetesen egy szimpla atirassal maradhattam volna hu a feladathoz, de akkor az openapi altal generaltba kellett volna modositanom. 

BACKEND: 
A feladatot hexagonal arhitekturaban oldottam meg.
Ez az arhitektura kicsit overkill ehez a feladathoz, de nekem tetszik. (atlathatosag, ujrafelhasznalas, konnyebb tesztelhetoseg)

A driver (adapter) itt az internal/httphandler
driven (adapter) internal/repositories

internal/core/services-ben vannak a query, command, event interakciok megvalositasa 
ezek felelosek a domain logic es a repo kezeleseert

A pkg dir-ben vannak a konnyen mashol is hasznalhato csomagok. (log,error,httpserver kiegeszitok)

internal/httphandler - felelos a request ellenorzesert, response gyartasert

Hiba kezelesre kulon packaget hasznalok, logolassal. Lehetoseg van bizonyos hibakat eljuttatni a responsba a hiba tipusa alapjan.
(pkg/errors/httperror, exterror)

------------------------------------------------------------------------------------------
FRONTEND:
about,home,post component
home - megnyitasa betolti az osszes url-hash part
post - itt lehet addolni 
backend api hivasokat az angluar httpclient vegzi a data services-ben

