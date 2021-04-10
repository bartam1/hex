import { Injectable } from '@angular/core';
import { HttpClient,HttpErrorResponse } from '@angular/common/http';

import {  throwError } from 'rxjs';
import { retry, catchError } from 'rxjs/operators';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class DataService {

  private API_SERVER = environment.apiUrl;

  constructor(private httpClient: HttpClient) { }
  
  handleError(error: HttpErrorResponse) {
    let errorMessage = 'Unknown error!';
    if (error.error instanceof ErrorEvent) {
      // Client-side errors
      errorMessage = `Error: ${error.error.message}`;
    } else {
      // Server-side errors
      errorMessage = `Error Code: ${error.status}\nMessage: ${error.message}`;
    }
    window.alert(errorMessage);
    return throwError(errorMessage);
  }
  
  public GetUrlsWidthHash(){
    let url = this.API_SERVER + "/urls";
    return this.httpClient.get(url).pipe(retry(3),catchError(this.handleError));
  }
  public PostMakeUrlHash(mu: string){
    console.log("pressed")
    console.log(mu)
    let url = this.API_SERVER + "/urls/new";
    let reqBody = "Url=" + mu;
    console.log(reqBody);
    return this.httpClient.post(url,{'Url': mu}).pipe(catchError(this.handleError));
  }
  public RemoveUrl(u: string) {
    let url = this.API_SERVER + "/url/del/" + u;
    return this.httpClient.delete(url).pipe(catchError(this.handleError));
  }
}
