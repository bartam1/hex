import { Component, OnInit, OnDestroy } from '@angular/core';
import { DataService } from '../data.service';
import {  takeUntil } from 'rxjs/operators'
import { Subject } from 'rxjs';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit , OnDestroy {
  
  UrlsWidthHash = []
  destroy$: Subject<boolean> = new Subject<boolean>();

  constructor(private dataService: DataService) { }

  ngOnInit(): void {
    this.dataService.GetUrlsWidthHash().pipe(takeUntil(this.destroy$)).subscribe((data: any[])=>{
      console.log(data);
      this.UrlsWidthHash = data;
    })  
  }
  ngOnDestroy() {
    this.destroy$.next(true);
    this.destroy$.unsubscribe();
  }

  removeUrl(u: string) {
    this.dataService.RemoveUrl(u).subscribe((data: any[])=>{
      console.log(data);
  
    })
  }
}
