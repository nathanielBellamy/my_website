import { Component, OnInit, inject } from '@angular/core';
import { GrooveJrStore } from './groove-jr.store';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { InfiniteScrollComponent } from '../../components/infinite-scroll/infinite-scroll.component';
import { CardComponent } from '../../components/card/card.component';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';
import { PageComponent } from '../../components/page/page.component';
import { encodeId } from '../../utils/id-encoder';
import { getSnippet } from '../../utils/snippet';

@Component({
  selector: 'app-groove-jr',
  standalone: true,
  imports: [CommonModule, InfiniteScrollComponent, CardComponent, ScrollFadeInDirective, PageComponent],
  templateUrl: './groove-jr.component.html',
})
export class GrooveJrComponent implements OnInit {
  protected readonly store = inject(GrooveJrStore);
  private readonly router = inject(Router);

  ngOnInit() {
    this.store.loadMore();
  }

  onScroll() {
    this.store.loadMore();
  }

  viewContent(id: string) {
    this.router.navigate(['/groovejr', encodeId(id)]);
  }

  getSnippet(content: string): string {
    return getSnippet(content);
  }
}
