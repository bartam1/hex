import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class DataService {

  private API_SERVER = "http://localhost:3000";

  constructor(private httpClient: HttpClient) { }
  public GetUrlsWidthHash(){
    let url = this.API_SERVER + "/urls";
    return this.httpClient.get(url);
  }
  // public PostMakeUrlHash(){
  //   let url = this.API_SERVER + "/url/new";
  //   return this.httpClient.post(url);
  // }
}
