import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import UploadFile from './components/UploadFile.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
   <UploadFile />
  </StrictMode>,
)
