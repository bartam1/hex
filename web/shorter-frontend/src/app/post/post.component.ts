import { Component, OnInit } from '@angular/core';
import { DataService } from '../data.service';


@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.css']
})
export class PostComponent implements OnInit {
  urlstr: string
  constructor(private dataService: DataService) { }

  ngOnInit(): void {
  }
  addUrl() {
    this.dataService.PostMakeUrlHash(this.urlstr).subscribe((data: any[])=>{
      console.log(data);
  
    })
  }
}
    