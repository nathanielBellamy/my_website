import { ComponentFixture, TestBed } from '@angular/core/testing';
import { EditHomeContentComponent } from './edit-home-content.component';
import { HomeService } from '../../services/home.service';
import { ActivatedRoute, Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { of } from 'rxjs';
import { HomeContent } from '../../models/data-models';

describe('EditHomeContentComponent', () => {
  let component: EditHomeContentComponent;
  let fixture: ComponentFixture<EditHomeContentComponent>;
  let mockHomeService: Partial<HomeService>;
  let mockActivatedRoute: Partial<ActivatedRoute>;
  let mockRouter: Partial<Router>;

  const mockHomeContent: HomeContent = { id: '1', title: 'Existing Home', content: 'Existing Content' };

  beforeEach(async () => {
    mockHomeService = {
      getHomeContentById: jasmine.createSpy('getHomeContentById').and.returnValue(Promise.resolve(mockHomeContent)),
      updateHomeContent: jasmine.createSpy('updateHomeContent').and.returnValue(Promise.resolve(mockHomeContent)),
    };
    mockActivatedRoute = {
      paramMap: of(new Map([['id', '1']])),
    };
    mockRouter = {
      navigate: jasmine.createSpy('navigate'),
    };

    await TestBed.configureTestingModule({
      imports: [EditHomeContentComponent, FormsModule],
      providers: [
        { provide: HomeService, useValue: mockHomeService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: Router, useValue: mockRouter },
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(EditHomeContentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should fetch home content on init and populate form', async () => {
    // Wait for the promise to resolve
    await fixture.whenStable();

    expect(mockHomeService.getHomeContentById).toHaveBeenCalledWith('1');
    expect(component.homeContent()).toEqual(mockHomeContent);
  });

  it('should update home content and navigate on success', async () => {
    // Ensure content is loaded first
    await fixture.whenStable();

    const updatedTitle = 'Updated Title';
    const updatedContent = 'Updated Content';
    component.homeContent.set({ ...mockHomeContent, title: updatedTitle, content: updatedContent });

    await component.updateContent();

    expect(mockHomeService.updateHomeContent).toHaveBeenCalledWith({ ...mockHomeContent, title: updatedTitle, content: updatedContent });
    expect(mockRouter.navigate).toHaveBeenCalledWith(['/home']);
  });

  it('should navigate back to list on goBack', () => {
    component.goBack();
    expect(mockRouter.navigate).toHaveBeenCalledWith(['/home']);
  });
});
