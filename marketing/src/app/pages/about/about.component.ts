import { Component, OnInit, inject } from '@angular/core';
import { AboutStore } from './about.store';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-about',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './about.component.html',
})
export class AboutComponent implements OnInit {
  readonly store = inject(AboutStore);

  ngOnInit() {
    this.store.loadContent();
  }
}
