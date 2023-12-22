import React from 'react';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import Paginator from './Paginator';

describe('<Paginator />', () => {
  test('it should mount', () => {
    render(<Paginator />);
    
    const paginator = screen.getByTestId('Paginator');

    expect(paginator).toBeInTheDocument();
  });
});