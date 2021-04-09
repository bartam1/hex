set -e

npm install
npm install -g @angular/cli@11.2.7
ng add @angular/material
ng serve --prod --host 0.0.0.0
