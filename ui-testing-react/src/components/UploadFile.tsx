import  { useState } from "react";
import { uploadFileTos3 } from "../../api"

const UploadFile = () => {
  const [file, setFile] = useState<File | null>(null);
  const [uploading, setUploading] = useState(false);
  const [error, setError] = useState<string | null>(null);

 const handleUpload = async () => {
    if (!file) return;
    setUploading(true);
    setError(null);
    try {
        const response=await uploadFileTos3(file);
      if (response.status === 200) {
        alert("File uploaded successfully!");
      } else {
        alert("File upload failed.");
      }
    } catch (err) {
      console.error("Error uploading file:", err);
      setError("File upload failed. Please try again.");
    } finally {
      setUploading(false);
      setFile(null);
    }
  };

  return (
    <div className="p-6 max-w-md mx-auto bg-white rounded-xl shadow-md space-y-4 border border-gray-200">
      <h3 className="text-xl font-semibold text-gray-800">Upload to S3</h3>
      <input
        type="file"
        onChange={(e) => {
          const selectedFile = e.target.files ? e.target.files[0] : null;
          setFile(selectedFile);
        }}
        className="block w-full text-sm text-gray-600 file:mr-4 file:py-2 file:px-4 file:border-0
          file:text-sm file:font-semibold file:bg-blue-50 file:text-blue-700 hover:file:bg-blue-100"
      />
      <button
        onClick={handleUpload}
        disabled={!file || uploading}
        className={`w-full py-2 px-4 text-white font-semibold rounded-lg shadow-md transition duration-200 ${
          uploading || !file
            ? "bg-gray-400 cursor-not-allowed"
            : "bg-blue-600 hover:bg-blue-700"
        }`}
      >
        {uploading ? "Uploading..." : "Upload"}
      </button>
    </div>
  );
};

export default UploadFile;
