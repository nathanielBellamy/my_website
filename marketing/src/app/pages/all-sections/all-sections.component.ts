import { Component, OnInit, AfterViewInit, OnDestroy, inject, NgZone } from '@angular/core';
import { CommonModule, Location } from '@angular/common';
import { Router, NavigationEnd } from '@angular/router';
import { filter, debounceTime, throttleTime } from 'rxjs/operators';
import { fromEvent, Subscription } from 'rxjs';
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
  private scrollSubscription?: Subscription;
  private scrollEndSubscription?: Subscription;
  private isAutoScrolling = false;

  ngOnInit() {
    this.router.events.pipe(
      filter(event => event instanceof NavigationEnd)
    ).subscribe(() => {
      this.scrollToSection();
    });

    // Run scroll listener outside Angular zone to prevent excessive change detection cycles
    this.ngZone.runOutsideAngular(() => {
      // Throttle scroll events for performance
      this.scrollSubscription = fromEvent(window, 'scroll')
        .pipe(throttleTime(50))
        .subscribe(() => {
          this.onScroll();
        });

      // Detect end of scrolling to reset isAutoScrolling flag
      this.scrollEndSubscription = fromEvent(window, 'scroll')
        .pipe(debounceTime(150))
        .subscribe(() => {
          this.isAutoScrolling = false;
        });
    });
  }

  ngAfterViewInit() {
    // Small delay to ensure DOM is ready
    setTimeout(() => this.scrollToSection(), 100);
  }

  ngOnDestroy() {
    this.scrollSubscription?.unsubscribe();
    this.scrollEndSubscription?.unsubscribe();
  }

  private onScroll() {
    if (this.isAutoScrolling) return;

    const sections = [
      { id: 'home', path: '' },
      { id: 'about', path: 'about' },
      { id: 'groovejr', path: 'groovejr' },
      { id: 'blog', path: 'blog' },
    ];

    const scrollPosition = window.scrollY + (window.innerHeight / 2); // Active point center of the screen

    let activeSection = sections[0];

    for (const section of sections) {
      const element = document.getElementById(section.id);
      if (element) {
        const offsetTop = element.offsetTop;
        const offsetBottom = offsetTop + element.offsetHeight;

        if (scrollPosition >= offsetTop && scrollPosition < offsetBottom) {
          activeSection = section;
          break;
        }
      }
    }

    const currentPath = this.location.path().replace(/^\//, ''); // Remove leading slash
    // Handle the root path case which might return ''
    const normalizedCurrentPath = currentPath === '' ? '' : currentPath; 
    
    if (normalizedCurrentPath !== activeSection.path) {
      this.ngZone.run(() => {
        this.location.replaceState(activeSection.path);
      });
    }
  }

  private scrollToSection() {
    this.isAutoScrolling = true;
    const url = this.router.url;
    let sectionId = 'home';
    
    // Simple matching logic
    if (url.includes('about')) sectionId = 'about';
    else if (url.includes('groovejr')) sectionId = 'groovejr';
    else if (url.includes('blog')) sectionId = 'blog';
    
    const element = document.getElementById(sectionId);
    if (element) {
      element.scrollIntoView({ behavior: 'smooth', block: 'start' });
    } else {
        // fallback to top if element not found (e.g. home)
        window.scrollTo({ top: 0, behavior: 'smooth' });
    }
  }
}
