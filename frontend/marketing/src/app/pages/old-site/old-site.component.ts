import { Component, AfterViewInit, OnDestroy, ElementRef, ViewChild } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PageComponent } from '../../components/page/page.component';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';
import { environment } from '../../../environments/environment';
import init, { MagicSquare } from '../../../../pkg/src_rust.js';

@Component({
  selector: 'app-old-site',
  standalone: true,
  imports: [CommonModule, PageComponent, ScrollFadeInDirective],
  templateUrl: './old-site.component.html'
})
export class OldSiteComponent implements AfterViewInit, OnDestroy {
  @ViewChild('appMain', { static: false }) appMainRef?: ElementRef;
  private magicSquarePromise: Promise<any> | null = null;

  readonly oldSiteUrl = environment.BASE_URL_OLD_SITE;

  async ngAfterViewInit(): Promise<void> {
    try {
      // Load and initialize the WASM module
      await init('/wasm/src_rust_bg.wasm');
      
      // Initializing the app with default settings, default presets, and false for touch_screen
      this.magicSquarePromise = MagicSquare.run(null, null, false);
      await this.magicSquarePromise;
      
    } catch (e) {
      console.error('Failed to initialize MagicSquare WASM module:', e);
    }
  }

  ngOnDestroy(): void {
    // Rust code listens for the 'destroymswasm' event on the 'app_main' element
    // to stop requestAnimationFrame loop
    const appMain = this.appMainRef?.nativeElement || document.getElementById('app_main');
    if (appMain) {
      appMain.dispatchEvent(new Event('destroymswasm'));
    }
  }
}
