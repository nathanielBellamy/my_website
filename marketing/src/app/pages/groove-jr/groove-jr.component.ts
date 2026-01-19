import { Component, OnInit, inject } from '@angular/core';
import { GrooveJrStore } from './groove-jr.store';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-groove-jr',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './groove-jr.component.html',
})
export class GrooveJrComponent implements OnInit {
  readonly store = inject(GrooveJrStore);

  ngOnInit() {
    this.store.loadContent();
  }
}
