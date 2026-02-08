import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HomeContentListComponent } from './home-content-list.component';
import { HomeService } from '../../services/home.service';
import { of } from 'rxjs';
import { HomeContent } from '../../models/data-models';

describe('HomeContentListComponent', () => {
  let component: HomeContentListComponent;
  let fixture: ComponentFixture<HomeContentListComponent>;
  let mockHomeService: Partial<HomeService>;

  const mockHomeContent: HomeContent[] = [
    { id: '1', title: 'Home 1', content: 'Content 1' },
    { id: '2', title: 'Home 2', content: 'Content 2' },
  ];

  beforeEach(async () => {
    mockHomeService = {
      getAllHomeContent: jasmine.createSpy('getAllHomeContent').and.returnValue(Promise.resolve(mockHomeContent)),
      deleteHomeContent: jasmine.createSpy('deleteHomeContent').and.returnValue(Promise.resolve()),
    };

    await TestBed.configureTestingModule({
      imports: [HomeContentListComponent],
      providers: [{ provide: HomeService, useValue: mockHomeService }]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HomeContentListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should fetch home content on ngOnInit', async () => {
    // Wait for the promise to resolve
    await fixture.whenStable();

    expect(mockHomeService.getAllHomeContent).toHaveBeenCalled();
    expect(component.homeContent()).toEqual(mockHomeContent);
  });

  it('should delete home content and refresh the list', async () => {
    // Initial fetch
    await fixture.whenStable();
    expect(component.homeContent()).toEqual(mockHomeContent);

    // Call delete
    component.deleteContent('1');

    // Wait for delete promise to resolve and then the refresh promise
    await fixture.whenStable();

    expect(mockHomeService.deleteHomeContent).toHaveBeenCalledWith('1');
    expect(mockHomeService.getAllHomeContent).toHaveBeenCalledTimes(2); // Initial fetch + refresh
  });
});
