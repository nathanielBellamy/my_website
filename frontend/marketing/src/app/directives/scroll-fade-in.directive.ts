import { Directive, ElementRef, OnInit, OnDestroy, Input, Renderer2, inject } from '@angular/core';

@Directive({
  selector: '[appScrollFadeIn]',
  standalone: true,
})
export class ScrollFadeInDirective implements OnInit, OnDestroy {
  private readonly el = inject(ElementRef);
  private readonly renderer = inject(Renderer2);
  private observer?: IntersectionObserver;

  @Input() delay: number = 0;
  @Input() threshold: number = 0;
  @Input() duration: number = 700;

  ngOnInit() {
    this.setupInitialState();
    this.setupObserver();
  }

  ngOnDestroy() {
    this.observer?.disconnect();
  }

  private setupInitialState() {
    this.renderer.addClass(this.el.nativeElement, 'opacity-0');
    this.renderer.addClass(this.el.nativeElement, 'translate-y-12');
    this.renderer.addClass(this.el.nativeElement, 'transition-all');
    this.renderer.setStyle(this.el.nativeElement, 'transition-duration', `${this.duration}ms`);
    this.renderer.setStyle(this.el.nativeElement, 'transition-timing-function', 'cubic-bezier(0.16, 1, 0.3, 1)'); // Custom ease-out
    
    if (this.delay > 0) {
      this.renderer.setStyle(this.el.nativeElement, 'transition-delay', `${this.delay}ms`);
    }
  }

  private setupObserver() {
    const options = {
      root: null,
      threshold: this.threshold,
    };

    this.observer = new IntersectionObserver((entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          this.renderer.removeClass(this.el.nativeElement, 'opacity-0');
          this.renderer.removeClass(this.el.nativeElement, 'translate-y-12');
          this.renderer.addClass(this.el.nativeElement, 'opacity-100');
          this.renderer.addClass(this.el.nativeElement, 'translate-y-0');
          this.observer?.unobserve(this.el.nativeElement);
        }
      });
    }, options);

    this.observer.observe(this.el.nativeElement);
  }
}
