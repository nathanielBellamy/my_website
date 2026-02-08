import { ComponentFixture, TestBed } from '@angular/core/testing';
import { CreateHomeContentComponent } from './create-home-content.component';
import { HomeService } from '../../services/home.service';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { HomeContent } from '../../models/data-models';

describe('CreateHomeContentComponent', () => {
  let component: CreateHomeContentComponent;
  let fixture: ComponentFixture<CreateHomeContentComponent>;
  let mockHomeService: Partial<HomeService>;
  let mockRouter: Partial<Router>;

  beforeEach(async () => {
    mockHomeService = {
      createHomeContent: jasmine.createSpy('createHomeContent').and.returnValue(Promise.resolve({ id: '1', title: 'New', content: 'Test' })),
    };
    mockRouter = {
      navigate: jasmine.createSpy('navigate'),
    };

    await TestBed.configureTestingModule({
      imports: [CreateHomeContentComponent, FormsModule],
      providers: [
        { provide: HomeService, useValue: mockHomeService },
        { provide: Router, useValue: mockRouter },
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(CreateHomeContentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should create home content and navigate on success', async () => {
    const newContent: HomeContent = { id: '', title: 'Test Title', content: 'Test Content' };
    component.homeContent = { ...newContent };

    await component.createContent();

    expect(mockHomeService.createHomeContent).toHaveBeenCalledWith(newContent);
    expect(mockRouter.navigate).toHaveBeenCalledWith(['/home']);
  });

  it('should navigate back to list on goBack', () => {
    component.goBack();
    expect(mockRouter.navigate).toHaveBeenCalledWith(['/home']);
  });
});
