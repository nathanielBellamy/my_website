import { ComponentFixture, TestBed } from '@angular/core/testing';
import { EditAboutContentComponent } from './edit-about-content.component';
import { AboutService } from '../../services/about.service';
import { ActivatedRoute, Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { of } from 'rxjs';
import { AboutContent } from '../../models/data-models';

describe('EditAboutContentComponent', () => {
  let component: EditAboutContentComponent;
  let fixture: ComponentFixture<EditAboutContentComponent>;
  let mockAboutService: Partial<AboutService>;
  let mockActivatedRoute: Partial<ActivatedRoute>;
  let mockRouter: Partial<Router>;

  const mockAboutContent: AboutContent = { id: '1', title: 'Existing About', content: 'Existing Content' };

  beforeEach(async () => {
    mockAboutService = {
      getAboutContentById: jasmine.createSpy('getAboutContentById').and.returnValue(Promise.resolve(mockAboutContent)),
      updateAboutContent: jasmine.createSpy('updateAboutContent').and.returnValue(Promise.resolve(mockAboutContent)),
    };
    mockActivatedRoute = {
      paramMap: of(new Map([['id', '1']])),
    };
    mockRouter = {
      navigate: jasmine.createSpy('navigate'),
    };

    await TestBed.configureTestingModule({
      imports: [EditAboutContentComponent, FormsModule],
      providers: [
        { provide: AboutService, useValue: mockAboutService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: Router, useValue: mockRouter },
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(EditAboutContentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should fetch about content on init and populate form', async () => {
    await fixture.whenStable();

    expect(mockAboutService.getAboutContentById).toHaveBeenCalledWith('1');
    expect(component.aboutContent()).toEqual(mockAboutContent);
  });

  it('should update about content and navigate on success', async () => {
    await fixture.whenStable();

    const updatedTitle = 'Updated Title';
    const updatedContent = 'Updated Content';
    component.aboutContent.set({ ...mockAboutContent, title: updatedTitle, content: updatedContent });

    await component.updateContent();

    expect(mockAboutService.updateAboutContent).toHaveBeenCalledWith({ ...mockAboutContent, title: updatedTitle, content: updatedContent });
    expect(mockRouter.navigate).toHaveBeenCalledWith(['/about']);
  });

  it('should navigate back to list on goBack', () => {
    component.goBack();
    expect(mockRouter.navigate).toHaveBeenCalledWith(['/about']);
  });
});
