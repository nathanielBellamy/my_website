import { Component, OnInit, inject } from '@angular/core';
import { CardComponent } from '../../components/card/card.component';
import { HomeStore } from './home.store';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [CardComponent, CommonModule],
  templateUrl: './home.component.html',
})
export class HomeComponent implements OnInit {
  readonly store = inject(HomeStore);

  ngOnInit() {
    this.store.loadContent();
  }
}
