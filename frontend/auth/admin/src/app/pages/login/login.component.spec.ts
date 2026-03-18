import { render, screen, fireEvent, waitFor } from '@testing-library/angular';
import { LoginComponent } from './login.component';
import { AuthService } from '../../services/auth.service';
import { provideHttpClient } from '@angular/common/http';

describe('LoginComponent', () => {
  const mockAuthService = {
    getChallenge: jest.fn(),
    validatePassword: jest.fn(),
    requestOtp: jest.fn(),
    verifyOtp: jest.fn(),
  };

  async function renderLogin() {
    return render(LoginComponent, {
      providers: [
        provideHttpClient(),
        { provide: AuthService, useValue: mockAuthService },
      ],
    });
  }

  beforeEach(() => {
    jest.resetAllMocks();
  });

  describe('password step', () => {
    it('should render the password form by default', async () => {
      await renderLogin();

      expect(screen.getByText('Admin Login')).toBeInTheDocument();
      expect(screen.getByLabelText('Admin Password')).toBeInTheDocument();
      expect(screen.getByRole('button', { name: 'Next' })).toBeInTheDocument();
    });

    it('should disable the Next button when password is empty', async () => {
      await renderLogin();

      expect(screen.getByRole('button', { name: 'Next' })).toBeDisabled();
    });

    it('should enable the Next button when password is entered', async () => {
      await renderLogin();

      await fireEvent.input(screen.getByLabelText('Admin Password'), {
        target: { value: 'mypassword' },
      });

      expect(screen.getByRole('button', { name: 'Next' })).toBeEnabled();
    });

    it('should call getChallenge and validatePassword on submit', async () => {
      mockAuthService.getChallenge.mockResolvedValue('test-challenge');
      mockAuthService.validatePassword.mockResolvedValue(undefined);

      await renderLogin();

      await fireEvent.input(screen.getByLabelText('Admin Password'), {
        target: { value: 'secret' },
      });
      await fireEvent.click(screen.getByRole('button', { name: 'Next' }));

      await waitFor(() => {
        expect(mockAuthService.getChallenge).toHaveBeenCalled();
        expect(mockAuthService.validatePassword).toHaveBeenCalledWith(
          expect.any(String)
        );
      });
    });

    it('should advance to the request step on successful password', async () => {
      mockAuthService.getChallenge.mockResolvedValue('challenge');
      mockAuthService.validatePassword.mockResolvedValue(undefined);

      await renderLogin();

      await fireEvent.input(screen.getByLabelText('Admin Password'), {
        target: { value: 'correct' },
      });
      await fireEvent.click(screen.getByRole('button', { name: 'Next' }));

      await waitFor(() => {
        expect(screen.getByText('Send Admin OTP')).toBeInTheDocument();
      });
    });

    it('should display an error on invalid password', async () => {
      mockAuthService.getChallenge.mockImplementation(
        () => new Promise((_, reject) => setTimeout(() => reject(new Error('fail')), 0))
      );

      // Suppress the expected console.error from the component
      const spy = jest.spyOn(console, 'error').mockImplementation(() => {});

      await renderLogin();

      await fireEvent.input(screen.getByLabelText('Admin Password'), {
        target: { value: 'wrong' },
      });
      await fireEvent.click(screen.getByRole('button', { name: 'Next' }));

      await waitFor(() => {
        expect(screen.getByText('Invalid Password')).toBeInTheDocument();
      });
      spy.mockRestore();
    });

    it('should show Validating... while loading', async () => {
      let resolveChallenge!: (v: string) => void;
      mockAuthService.getChallenge.mockReturnValue(
        new Promise<string>((r) => { resolveChallenge = r; })
      );

      await renderLogin();

      await fireEvent.input(screen.getByLabelText('Admin Password'), {
        target: { value: 'pw' },
      });
      await fireEvent.click(screen.getByRole('button', { name: 'Next' }));

      await waitFor(() => {
        expect(screen.getByText('Validating...')).toBeInTheDocument();
      });

      resolveChallenge('c');
    });
  });

  describe('request OTP step', () => {
    async function advanceToRequestStep() {
      mockAuthService.getChallenge.mockResolvedValue('challenge');
      mockAuthService.validatePassword.mockResolvedValue(undefined);

      await renderLogin();

      await fireEvent.input(screen.getByLabelText('Admin Password'), {
        target: { value: 'pass' },
      });
      await fireEvent.click(screen.getByRole('button', { name: 'Next' }));

      await waitFor(() => {
        expect(screen.getByText('Send Admin OTP')).toBeInTheDocument();
      });
    }

    it('should call requestOtp when the button is clicked', async () => {
      mockAuthService.requestOtp.mockResolvedValue(undefined);
      await advanceToRequestStep();

      await fireEvent.click(screen.getByText('Send Admin OTP'));

      await waitFor(() => {
        expect(mockAuthService.requestOtp).toHaveBeenCalled();
      });
    });

    it('should advance to verify step after OTP is sent', async () => {
      mockAuthService.requestOtp.mockResolvedValue(undefined);
      await advanceToRequestStep();

      await fireEvent.click(screen.getByText('Send Admin OTP'));

      await waitFor(() => {
        expect(screen.getByLabelText('One-Time Passcode')).toBeInTheDocument();
      });
    });

    it('should display an error if OTP request fails', async () => {
      mockAuthService.requestOtp.mockImplementation(
        () => new Promise((_, reject) => setTimeout(() => reject(new Error('Email failed')), 0))
      );
      await advanceToRequestStep();

      await fireEvent.click(screen.getByText('Send Admin OTP'));

      await waitFor(() => {
        expect(screen.getByText('Email failed')).toBeInTheDocument();
      });
    });
  });

  describe('verify OTP step', () => {
    async function advanceToVerifyStep() {
      mockAuthService.getChallenge.mockResolvedValue('challenge');
      mockAuthService.validatePassword.mockResolvedValue(undefined);
      mockAuthService.requestOtp.mockResolvedValue(undefined);

      await renderLogin();

      await fireEvent.input(screen.getByLabelText('Admin Password'), {
        target: { value: 'pass' },
      });
      await fireEvent.click(screen.getByRole('button', { name: 'Next' }));

      await waitFor(() => {
        expect(screen.getByText('Send Admin OTP')).toBeInTheDocument();
      });

      await fireEvent.click(screen.getByText('Send Admin OTP'));

      await waitFor(() => {
        expect(screen.getByLabelText('One-Time Passcode')).toBeInTheDocument();
      });
    }

    it('should render the OTP input and Login button', async () => {
      await advanceToVerifyStep();

      expect(screen.getByLabelText('One-Time Passcode')).toBeInTheDocument();
      expect(screen.getByRole('button', { name: 'Login' })).toBeInTheDocument();
    });

    it('should disable Login when OTP is empty', async () => {
      await advanceToVerifyStep();

      expect(screen.getByRole('button', { name: 'Login' })).toBeDisabled();
    });

    it('should call verifyOtp on submit', async () => {
      mockAuthService.verifyOtp.mockResolvedValue(undefined);
      await advanceToVerifyStep();

      await fireEvent.input(screen.getByLabelText('One-Time Passcode'), {
        target: { value: '123456' },
      });
      await fireEvent.click(screen.getByRole('button', { name: 'Login' }));

      await waitFor(() => {
        expect(mockAuthService.verifyOtp).toHaveBeenCalledWith('123456');
      });
    });

    it('should display error on invalid OTP', async () => {
      mockAuthService.verifyOtp.mockImplementation(
        () => new Promise((_, reject) => setTimeout(() => reject(new Error('Invalid OTP')), 0))
      );
      await advanceToVerifyStep();

      await fireEvent.input(screen.getByLabelText('One-Time Passcode'), {
        target: { value: 'wrong' },
      });
      await fireEvent.click(screen.getByRole('button', { name: 'Login' }));

      await waitFor(() => {
        expect(screen.getByText('Invalid OTP')).toBeInTheDocument();
      });
    });

    it('should navigate back to OTP request step via Back button', async () => {
      await advanceToVerifyStep();

      await fireEvent.click(screen.getByText('Back'));

      await waitFor(() => {
        expect(screen.getByText('Send Admin OTP')).toBeInTheDocument();
      });
    });
  });
});
