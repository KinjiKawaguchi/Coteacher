import React, { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

type EmailInputProps = {
  handleSubmit: (email: string) => void; // または Promise<void>
};

const EmailInput = ({ handleSubmit }: EmailInputProps) => {
  const [email, setEmail] = useState('');

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(event.currentTarget.value);
  };

  const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter' && email === event.currentTarget.value) {
      handleSubmit(email);
    }
  };

  const handleEmailSubmit = () => {
    handleSubmit(email);
  };

  return (
    <>
      <Input
        type="email"
        value={email}
        placeholder="example@example.test"
        onChange={handleChange}
        onKeyDown={handleKeyDown}
      />
      <Button onClick={handleEmailSubmit}>ログインリンクを送信</Button>
    </>
  );
};

export default EmailInput;
