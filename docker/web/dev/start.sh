set -e

npm install
npm install -g @angular/cli@11.2.7
ng add @angular/material
envsubst < src/assets/env.template.js > src/assets/env.js
ng serve --host web
