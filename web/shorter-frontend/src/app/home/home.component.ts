import { Component, OnInit } from '@angular/core';
import { DataService } from '../data.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  UrlsWidthHash = []

  constructor(private dataService: DataService) { }

  ngOnInit(): void {
    this.dataService.GetUrlsWidthHash().subscribe((data: any[])=>{
      console.log(data);
      this.UrlsWidthHash = data;
    })  
  }

}
