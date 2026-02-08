import { ComponentFixture, TestBed } from '@angular/core/testing';
import { CreateAboutContentComponent } from './create-about-content.component';
import { AboutService } from '../../services/about.service';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { AboutContent } from '../../models/data-models';

describe('CreateAboutContentComponent', () => {
  let component: CreateAboutContentComponent;
  let fixture: ComponentFixture<CreateAboutContentComponent>;
  let mockAboutService: Partial<AboutService>;
  let mockRouter: Partial<Router>;

  beforeEach(async () => {
    mockAboutService = {
      createAboutContent: jasmine.createSpy('createAboutContent').and.returnValue(Promise.resolve({ id: '1', title: 'New', content: 'Test' })),
    };
    mockRouter = {
      navigate: jasmine.createSpy('navigate'),
    };

    await TestBed.configureTestingModule({
      imports: [CreateAboutContentComponent, FormsModule],
      providers: [
        { provide: AboutService, useValue: mockAboutService },
        { provide: Router, useValue: mockRouter },
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAboutContentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should create about content and navigate on success', async () => {
    const newContent: AboutContent = { id: '', title: 'Test Title', content: 'Test Content' };
    component.aboutContent = { ...newContent };

    await component.createContent();

    expect(mockAboutService.createAboutContent).toHaveBeenCalledWith(newContent);
    expect(mockRouter.navigate).toHaveBeenCalledWith(['/about']);
  });

  it('should navigate back to list on goBack', () => {
    component.goBack();
    expect(mockRouter.navigate).toHaveBeenCalledWith(['/about']);
  });
});
