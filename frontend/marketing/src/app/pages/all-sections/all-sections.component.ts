import { Component, OnInit, AfterViewInit, OnDestroy, inject, NgZone } from '@angular/core';
import { CommonModule, Location } from '@angular/common';
import { Router, NavigationEnd } from '@angular/router';
import { filter } from 'rxjs/operators';
import { Subscription } from 'rxjs';
import { HomeComponent } from '../home/home.component';
import { AboutComponent } from '../about/about.component';
import { GrooveJrComponent } from '../groove-jr/groove-jr.component';
import { BlogComponent } from '../blog/blog.component';

@Component({
  selector: 'app-all-sections',
  standalone: true,
  imports: [CommonModule, HomeComponent, AboutComponent, GrooveJrComponent, BlogComponent],
  templateUrl: './all-sections.component.html',
})
export class AllSectionsComponent implements OnInit, AfterViewInit, OnDestroy {
  private readonly router = inject(Router);
  private readonly location = inject(Location);
  private readonly ngZone = inject(NgZone);
  private routerSubscription?: Subscription;
  private observer?: IntersectionObserver;
  private isAutoScrolling = false;
  private lastNavTime = 0;

  ngOnInit() {
    this.routerSubscription = this.router.events.pipe(
      filter(event => event instanceof NavigationEnd)
    ).subscribe(() => {
      this.isAutoScrolling = true;
      this.lastNavTime = Date.now();
      
      // Small delay to allow scroll restoration to complete first
      setTimeout(() => {
        this.scrollToSection();
      }, 100);

      // Allow scroll tracking after a delay
      setTimeout(() => {
        this.isAutoScrolling = false;
      }, 1000);
    });
  }

  ngAfterViewInit() {
    this.setupIntersectionObserver();
    // Small delay to ensure DOM is ready
    setTimeout(() => this.scrollToSection(), 100);
  }

  ngOnDestroy() {
    this.routerSubscription?.unsubscribe();
    this.observer?.disconnect();
  }

  private setupIntersectionObserver() {
    const options = {
      root: null,
      rootMargin: '-20% 0px -70% 0px', // Trigger when section is in the top portion of the viewport
      threshold: 0
    };

    this.observer = new IntersectionObserver((entries) => {
      if (this.isAutoScrolling || (Date.now() - this.lastNavTime < 1000)) return;

      entries.forEach(entry => {
        if (entry.isIntersecting) {
          const sectionId = entry.target.id;
          const path = sectionId === 'home' ? '' : sectionId;
          const currentPath = this.location.path().replace(/^\//, '');
          
          if (currentPath !== path) {
            this.ngZone.run(() => {
              this.location.replaceState(path);
            });
          }
        }
      });
    }, options);

    const sections = ['home', 'about', 'groovejr', 'blog'];
    sections.forEach(id => {
      const element = document.getElementById(id);
      if (element) {
        this.observer?.observe(element);
      }
    });
  }

  private scrollToSection() {
    const url = this.router.url;
    let sectionId = 'home';
    
    if (url.includes('about')) sectionId = 'about';
    else if (url.includes('groovejr')) sectionId = 'groovejr';
    else if (url.includes('blog')) sectionId = 'blog';
    
    const element = document.getElementById(sectionId);
    if (element) {
      element.scrollIntoView({ behavior: 'smooth', block: 'start' });
    } else {
        window.scrollTo({ top: 0, behavior: 'smooth' });
    }
  }
}
