import { Component, OnInit, inject } from '@angular/core';
import { BlogStore } from './blog.store';
import { CommonModule } from '@angular/common';
import { CardComponent } from '../../components/card/card.component';

@Component({
  selector: 'app-blog',
  standalone: true,
  imports: [CommonModule, CardComponent],
  templateUrl: './blog.component.html',
})
export class BlogComponent implements OnInit {
  readonly store = inject(BlogStore);

  ngOnInit() {
    this.store.loadPosts();
  }
}
