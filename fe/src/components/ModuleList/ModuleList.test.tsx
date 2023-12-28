import React from 'react';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import ModuleList from './ModuleList';

describe('<ModuleList />', () => {
  test('it should mount', () => {
    render(<ModuleList />);
    
    const moduleList = screen.getByTestId('ModuleList');

    expect(moduleList).toBeInTheDocument();
  });
});