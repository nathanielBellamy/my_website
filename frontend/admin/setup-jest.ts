import { setupZoneTestEnv } from 'jest-preset-angular/setup-env/zone';
import '@testing-library/jest-dom';

setupZoneTestEnv();

// Mock IntersectionObserver
class IntersectionObserver {
  readonly root: Element | Document | null = null;
  readonly rootMargin: string = '';
  readonly thresholds: ReadonlyArray<number> = [];

  constructor(public callback: IntersectionObserverCallback, public options?: IntersectionObserverInit) {}

  observe(target: Element): void {
    // Immediately trigger the callback with isIntersecting=true to simulate visibility in tests
    this.callback([{
      isIntersecting: true,
      target: target,
      intersectionRatio: 1,
      boundingClientRect: target.getBoundingClientRect(),
      intersectionRect: target.getBoundingClientRect(),
      rootBounds: null,
      time: Date.now()
    }] as IntersectionObserverEntry[], this);
  }
  unobserve(target: Element): void {}
  disconnect(): void {}
  takeRecords(): IntersectionObserverEntry[] { return []; }
}

Object.defineProperty(window, 'IntersectionObserver', {
  writable: true,
  configurable: true,
  value: IntersectionObserver,
});

// Mock window.confirm
window.confirm = jest.fn(() => true);

// Mock ngx-markdown to avoid parsing issues with marked (ESM)
jest.mock('ngx-markdown', () => {
  const { Component } = require('@angular/core');
  return {
    MarkdownComponent: Component({
      selector: 'markdown',
      standalone: true,
      template: '<ng-content></ng-content>',
      inputs: [
        'data',
        'src',
        'disableSanitizer',
        'inline',
        'clipboard',
        'clipboardButtonComponent',
        'clipboardButtonTemplate',
        'emoji',
        'katex',
        'katexOptions',
        'mermaid',
        'mermaidOptions',
        'lineHighlight',
        'line',
        'lineOffset',
        'lineNumbers',
        'start',
        'commandLine',
        'filterOutput',
        'host',
        'prompt',
        'output',
        'user',
      ]
    })(class {}),
    provideMarkdown: () => [],
    MarkdownModule: {
      forRoot: () => ({
        ngModule: class {},
      })
    }
  };
});
