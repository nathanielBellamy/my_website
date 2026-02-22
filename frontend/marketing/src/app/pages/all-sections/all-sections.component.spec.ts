import { ComponentFixture, TestBed, fakeAsync, tick } from '@angular/core/testing';
import { AllSectionsComponent } from './all-sections.component';
import { Router, NavigationEnd, ActivatedRoute } from '@angular/router';
import { Location } from '@angular/common';
import { Title } from '@angular/platform-browser';
import { Subject, of } from 'rxjs';
import { HttpClientTestingModule } from '@angular/common/http/testing';

describe('AllSectionsComponent', () => {
  let component: AllSectionsComponent;
  let fixture: ComponentFixture<AllSectionsComponent>;
  let routerEventsSubject: Subject<any>;
  let routerMock: any;
  let locationMock: any;
  let titleMock: any;

  beforeEach(async () => {
    routerEventsSubject = new Subject<any>();
    routerMock = {
      events: routerEventsSubject.asObservable(),
      url: '/blog',
      createUrlTree: jest.fn(),
      serializeUrl: jest.fn(),
      navigate: jest.fn()
    };
    locationMock = {
      path: jest.fn().mockReturnValue('/blog'),
      replaceState: jest.fn()
    };
    titleMock = {
      setTitle: jest.fn()
    };

    // Mock IntersectionObserver
    (window as any).IntersectionObserver = jest.fn().mockImplementation((callback) => ({
      observe: jest.fn(),
      disconnect: jest.fn(),
      takeRecords: jest.fn()
    }));

    // Mock document.getElementById
    jest.spyOn(document, 'getElementById').mockImplementation((id) => {
      return {
        scrollTo: jest.fn(),
        scrollIntoView: jest.fn(),
        getBoundingClientRect: jest.fn().mockReturnValue({ top: 0, bottom: 0 }),
        id: id
      } as any;
    });

    await TestBed.configureTestingModule({
      imports: [AllSectionsComponent, HttpClientTestingModule],
      providers: [
        { provide: Router, useValue: routerMock },
        { provide: Location, useValue: locationMock },
        { provide: Title, useValue: titleMock },
        { 
          provide: ActivatedRoute, 
          useValue: { 
            snapshot: { paramMap: { get: () => null } },
            queryParams: of({}),
            params: of({})
          } 
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(AllSectionsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  it('should initialize correctly', () => {
    expect(component).toBeTruthy();
  });

  it('should handle navigation events and delay observer logic', fakeAsync(() => {
    // Trigger navigation event
    routerEventsSubject.next(new NavigationEnd(1, '/blog', '/blog'));
    
    // Check if isAutoScrolling is true immediately
    expect((component as any).isAutoScrolling).toBe(true);
    
    // Fast forward 100ms (scrollToSection happens)
    tick(100);
    // isAutoScrolling should still be true
    expect((component as any).isAutoScrolling).toBe(true);

    // Fast forward to just before 2500ms
    tick(2399); 
    expect((component as any).isAutoScrolling).toBe(true);

    // Fast forward to 2500ms
    tick(1);
    
    // isAutoScrolling should become false
    expect((component as any).isAutoScrolling).toBe(false);
  }));
});
