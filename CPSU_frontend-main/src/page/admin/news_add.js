import React, { useState, useCallback } from "react";
import Navbar from "../../component/navbar";
import Headers from "../../component/header";
import Footer from "../../component/footer";
import Menu from "../../component/menu";
import { CKEditor } from "@ckeditor/ckeditor5-react";
import ClassicEditor from "@ckeditor/ckeditor5-build-classic";
import "../../css/admin/news_add.css";
import axios from "axios";
import { useNavigate } from "react-router-dom";

// Cropper import
import Cropper from "react-easy-crop";
import Button from "@mui/material/Button";
// import { getOrientation } from "get-orientation/browser";
import { makeStyles } from "@mui/styles";

const useStyles = makeStyles({
  cropContainer: {
    position: "relative",
    width: "100%",
    height: 400,
    background: "#333",
  },
  controls: {
    display: "flex",
    flexDirection: "column",
    alignItems: "stretch",
    marginTop: 16,
  },
  sliderContainer: {
    display: "flex",
    flex: "1",
    alignItems: "center",
  },
});

const Add_News = () => {
  const classes = useStyles();
  const navigate = useNavigate();

  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [category, setCategory] = useState("1");
  const [urlDetail, setUrlDetail] = useState("");
  const [fileImage, setFileImage] = useState(null);

  const [newsImages, setNewsImages] = useState([]); // เก็บไฟล์หลายรูป
  const [newsPreviewUrls, setNewsPreviewUrls] = useState([]); // เก็บ preview หลายรูป

  const [imageSrc, setImageSrc] = useState(null);
  const [crop, setCrop] = useState({ x: 0, y: 0 });
  const [zoom, setZoom] = useState(1);
  const [croppedAreaPixels, setCroppedAreaPixels] = useState(null);

  const onCropComplete = useCallback((croppedArea, croppedAreaPixels) => {
    setCroppedAreaPixels(croppedAreaPixels);
  }, []);

  const onNewsImagesChange = (e) => {
    if (e.target.files && e.target.files.length > 0) {
      const files = Array.from(e.target.files);
      setNewsImages(files);

      // สร้าง preview urls
      const previews = files.map((file) => URL.createObjectURL(file));
      setNewsPreviewUrls(previews);
    }
  };

  const [previewUrl, setPreviewUrl] = useState(null); // สำหรับแสดง preview

  const showCroppedImage = useCallback(async () => {
    try {
      // crop เป็น blob
      const blob = await getCroppedImg(imageSrc, croppedAreaPixels);

      // สร้าง File
      const file = new File([blob], "cover.jpg", { type: blob.type });

      // เก็บลง state
      setFileImage(file);

      // แปลง blob เป็น URL สำหรับ preview
      const preview = URL.createObjectURL(blob);
      setPreviewUrl(preview);

      // ปิด cropper
      setImageSrc(null);
    } catch (e) {
      console.error(e);
    }
  }, [imageSrc, croppedAreaPixels]);

  // อ่านไฟล์จาก input
  const onFileChange = async (e) => {
    if (e.target.files && e.target.files.length > 0) {
      const file = e.target.files[0];
      let imageDataUrl = await readFile(file);

      try {
        // const orientation = await getOrientation(file);
      } catch (e) {
        console.warn("failed to detect the orientation");
      }

      setImageSrc(imageDataUrl);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const newNews = new FormData();
    newNews.append("title", title);
    newNews.append("content", content);
    newNews.append("type_id", category);
    newNews.append("detail_url", urlDetail);

    // ✅ ส่ง cover_image แยก (หลัง crop แล้ว)
    if (fileImage) {
      newNews.append("cover_image", fileImage);
    }

    // ✅ ส่งรูปข่าวประกอบหลายรูป (ไม่รวมรูปหน้าปก)
    newsImages.forEach((file) => {
      newNews.append("images", file);
    });

    try {
      await axios.post("http://localhost:8080/api/v1/admin/news", newNews, {
        headers: { "Content-Type": "multipart/form-data" },
      });

      alert("เผยแพร่ข่าวสารสำเร็จ");
      navigate("/admin/news");
    } catch (error) {
      console.error("Error submitting news:", error);
      alert("ไม่สามารถเผยแพร่ข่าวสารได้");
    }
  };

  const removeImage = (index) => {
    const newImages = [...newsImages];
    const newPreviewUrls = [...newsPreviewUrls];
    newImages.splice(index, 1);
    newPreviewUrls.splice(index, 1);
    setNewsImages(newImages);
    setNewsPreviewUrls(newPreviewUrls);
  };

  return (
    <>
      <Headers />
      <Navbar />
      <div className="container text-center">
        <div className="row">
          <div className="col-sm-3">
            <Menu />
          </div>
          <div className="col-sm-9 text-start">
            <div>
              <h3 className="news-title">เพิ่มข่าวสาร</h3>
            </div>

            {imageSrc ? (
              <>
                <div className={classes.cropContainer}>
                  <Cropper
                    image={imageSrc}
                    crop={crop}
                    zoom={zoom}
                    aspect={4 / 3}
                    onCropChange={setCrop}
                    onCropComplete={onCropComplete}
                    onZoomChange={setZoom}
                  />
                </div>

                <div className={classes.controls}>
                  <Button
                    onClick={showCroppedImage}
                    variant="contained"
                    sx={{
                      backgroundColor: "#173390", // เปลี่ยนสีเป็นสีเขียว
                      margin: "50px", // เพิ่ม margin รอบ ๆ ปุ่ม
                      "&:hover": {
                        backgroundColor: "#0b1a4aff", // สีเมื่อ hover
                      },
                    }}
                  >
                    ตกลง
                  </Button>
                </div>
              </>
            ) : (
              <form onSubmit={handleSubmit}>
                <div className="mb-3">
                  <label className="form-label">หัวข้อข่าวสาร</label>
                  <input
                    className="form-control"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                  />
                </div>

                <div className="mb-3">
                  <label className="form-label">เนื้อหาข่าวสาร</label>
                  <CKEditor
                    editor={ClassicEditor}
                    data={content}
                    onChange={(event, editor) => setContent(editor.getData())}
                  />
                </div>

                <div className="mb-3">
                  <label className="form-label">ประเภทข่าวสาร</label>
                  <select
                    className="form-select"
                    value={category}
                    onChange={(e) => setCategory(e.target.value)}
                  >
                    <option value="1">ข่าวประชาสัมพันธ์</option>
                    <option value="2">ทุนการศึกษา</option>
                    <option value="3">รางวัลที่ได้รับ</option>
                    <option value="4">กิจกรรมของภาควิชา</option>
                  </select>
                </div>

                <div className="mb-3">
                  <label className="form-label">รายละเอียดข่าวสาร</label>
                  <input
                    className="form-control"
                    value={urlDetail}
                    onChange={(e) => setUrlDetail(e.target.value)}
                  />
                </div>

                <div className="mb-3">
                  <label className="form-label">รูปภาพหน้าปกข่าว</label>
                  <input
                    className="form-control"
                    type="file"
                    accept="image/*"
                    onChange={onFileChange}
                  />
                  {previewUrl && (
                    <div className="mt-3">
                      <p>Preview รูปหน้าปก:</p>
                      <img
                        src={previewUrl}
                        alt="preview"
                        style={{ maxWidth: "300px", borderRadius: "8px" }}
                      />
                    </div>
                  )}
                </div>

                <div className="mb-3">
                  <label className="form-label">รูปภาพข่าว</label>
                  <input
                    className="form-control"
                    type="file"
                    multiple
                    onChange={onNewsImagesChange}
                  />

                  {newsPreviewUrls.length > 0 && (
                    <div className="mt-3">
                      <p>Preview รูปภาพข่าว:</p>
                      <div
                        style={{
                          display: "flex",
                          flexWrap: "wrap",
                          gap: "10px",
                        }}
                      >
                        {newsPreviewUrls.map((url, index) => (
                          <div key={index} className="image-preview-wrapper">
                            <img
                              src={url}
                              alt={`preview-${index}`}
                              className="image-preview"
                            />
                            <button
                              type="button"
                              className="remove-image-btn"
                              onClick={() => removeImage(index)}
                            >
                              &times;
                            </button>
                          </div>
                        ))}
                      </div>
                    </div>
                  )}
                </div>

                <button
                  type="submit"
                  className="btn btn-primary"
                  id="btn-submit"
                >
                  เผยแพร่
                </button>
              </form>
            )}
          </div>
        </div>
      </div>
      <Footer />
    </>
  );
};

// ฟังก์ชันอ่านไฟล์เป็น dataURL
function readFile(file) {
  return new Promise((resolve) => {
    const reader = new FileReader();
    reader.addEventListener("load", () => resolve(reader.result), false);
    reader.readAsDataURL(file);
  });
}

// แปลง crop → blob
async function getCroppedImg(imageSrc, crop) {
  const image = await createImage(imageSrc);
  const canvas = document.createElement("canvas");
  const ctx = canvas.getContext("2d");

  canvas.width = crop.width;
  canvas.height = crop.height;

  ctx.drawImage(
    image,
    crop.x,
    crop.y,
    crop.width,
    crop.height,
    0,
    0,
    crop.width,
    crop.height
  );

  return new Promise((resolve) => {
    canvas.toBlob((blob) => resolve(blob), "image/jpeg");
  });
}

function createImage(url) {
  return new Promise((resolve, reject) => {
    const image = new Image();
    image.addEventListener("load", () => resolve(image));
    image.addEventListener("error", (error) => reject(error));
    image.setAttribute("crossOrigin", "anonymous");
    image.src = url;
  });
}

export default Add_News;
