import React, { useState, useEffect, useCallback } from "react";
import { useParams, useNavigate } from "react-router-dom";
import Navbar from "../../component/navbar";
import Headers from "../../component/header";
import Footer from "../../component/footer";
import Menu from "../../component/menu";
import { CKEditor } from "@ckeditor/ckeditor5-react";
import ClassicEditor from "@ckeditor/ckeditor5-build-classic";
import "../../css/admin/news_edit.css";
import axios from "axios";

// Cropper
import Cropper from "react-easy-crop";
import Button from "@mui/material/Button";
import { getOrientation } from "get-orientation/browser";

const Edit_News = () => {
  const { id } = useParams();
  const navigate = useNavigate();

  // state หลัก
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [category, setCategory] = useState("1");
  const [urlDetail, setUrlDetail] = useState("");

  // รูปภาพ
  const [fileImage, setFileImage] = useState(null); // cover ใหม่ ถ้ามี
  const [coverUrl, setCoverUrl] = useState(null);    // URL cover เดิม หรือ preview ใหม่

  const [newsImages, setNewsImages] = useState([]);        // File object ใหม่
  const [newsPreviewUrls, setNewsPreviewUrls] = useState([]); // URL preview (รวมใหม่ + เดิม)
  const [existingNewsUrls, setExistingNewsUrls] = useState([]); // URL ที่มาจาก backend เดิม

  // สำหรับ cropper
  const [imageSrc, setImageSrc] = useState(null);
  const [crop, setCrop] = useState({ x: 0, y: 0 });
  const [zoom, setZoom] = useState(1);
  const [croppedAreaPixels, setCroppedAreaPixels] = useState(null);

  useEffect(() => {
    if (!id) return;

    const fetchDetailNews = async () => {
      try {
        const res = await axios.get(
          `http://localhost:8080/api/v1/admin/news/${id}`
        );
        const data = res.data;

        setTitle(data.title || "");
        setContent(data.content || "");
        setCategory(String(data.type_id || 1));
        setUrlDetail(data.detail_url || "");

        if (data.images && data.images.length > 0) {
          // สมมติว่า images[0] เป็น cover
          const urls = data.images.map((img) => img.file_image);
          const cover = urls[0];
          const others = urls.slice(1);

          setCoverUrl(cover);
          setExistingNewsUrls(others);
          setNewsPreviewUrls(others);
        }
      } catch (err) {
        console.error("Error fetching news details", err);
      }
    };

    fetchDetailNews();
  }, [id]);

  const onNewsImagesChange = (e) => {
    if (e.target.files && e.target.files.length > 0) {
      const files = Array.from(e.target.files);
      // รวมกับไฟล์เดิม
      setNewsImages((prev) => [...prev, ...files]);
      const newPreviews = files.map((f) => URL.createObjectURL(f));
      setNewsPreviewUrls((prev) => [...prev, ...newPreviews]);
    }
  };

  const showCroppedImage = useCallback(async () => {
    try {
      const blob = await getCroppedImg(imageSrc, croppedAreaPixels);
      const file = new File([blob], "cover.jpg", { type: blob.type });
      setFileImage(file);
      setCoverUrl(URL.createObjectURL(blob));
      setImageSrc(null);
    } catch (e) {
      console.error(e);
    }
  }, [imageSrc, croppedAreaPixels]);

  const onCropComplete = useCallback((_, croppedAreaPixels) => {
    setCroppedAreaPixels(croppedAreaPixels);
  }, []);

  const onFileChange = async (e) => {
    if (e.target.files && e.target.files.length > 0) {
      const file = e.target.files[0];
      let imageDataUrl = await readFile(file);
      try {
        await getOrientation(file);
      } catch {
        console.warn("failed to detect orientation");
      }
      setImageSrc(imageDataUrl);
    }
  };

  const handleEdit = async (e) => {
    e.preventDefault();

    const formData = new FormData();
    formData.append("title", title);
    formData.append("content", content);
    formData.append("type_id", category);
    formData.append("detail_url", urlDetail);

    // cover
    if (fileImage) {
      formData.append("cover", fileImage);
    } else if (coverUrl) {
      // ถ้าไม่มี file ใหม่ แต่มี URL เดิม ให้ส่ง URL เดิม
      formData.append("existing_cover", coverUrl);
    }

    // news images
    if (newsImages.length > 0) {
      newsImages.forEach((file) => {
        formData.append("images", file);
      });
    }
    // ส่ง URL รูปข่าวที่ยังอยู่ (เดิม) ถ้าไม่มีไฟล์ใหม่
    if (existingNewsUrls.length > 0) {
      formData.append("existing_images", JSON.stringify(existingNewsUrls));
    }

    try {
      await axios.put(
        `http://localhost:8080/api/v1/admin/news/${id}`,
        formData,
        {
          headers: { "Content-Type": "multipart/form-data" },
        }
      );
      alert("แก้ไขข่าวสารสำเร็จ");
      navigate(`/admin/news/${id}`);
    } catch (err) {
      console.error("Error submitting news", err);
      alert("ไม่สามารถแก้ไขข่าวสารได้");
    }
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
            <p className="header-news-p">แก้ไขข่าวสาร</p>

            {imageSrc ? (
              <>
                <div
                  style={{
                    position: "relative",
                    width: "100%",
                    height: 400,
                    background: "#333",
                  }}
                >
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
                <div style={{ marginTop: 20, textAlign: "center" }}>
                  <Button
                    onClick={showCroppedImage}
                    variant="contained"
                    sx={{
                      backgroundColor: "#38419d",
                      "&:hover": { backgroundColor: "#2c327cff" },
                    }}
                  >
                    ตกลง
                  </Button>
                </div>
              </>
            ) : (
              <form onSubmit={handleEdit}>
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
                    onChange={(e, editor) => setContent(editor.getData())}
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
                  {coverUrl && (
                    <div
                      className="mt-3"
                      style={{ position: "relative", display: "inline-block" }}
                    >
                      <p>Preview รูปหน้าปก:</p>
                      <img
                        src={coverUrl}
                        alt="cover-preview"
                        style={{ maxWidth: "300px", borderRadius: "8px" }}
                      />
                      <button
                        type="button"
                        onClick={() => {
                          setCoverUrl(null);
                          setFileImage(null);
                        }}
                        style={{
                          position: "absolute",
                          top: "5px",
                          right: "5px",
                          background: "rgba(0,0,0,0.4)",
                          color: "white",
                          border: "none",
                          borderRadius: "50%",
                          width: "24px",
                          height: "24px",
                          fontSize: "18px",
                          cursor: "pointer",
                        }}
                        title="ลบรูปหน้าปก"
                      >
                        &times;
                      </button>
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

                  {(newsPreviewUrls.length > 0) && (
                    <div
                      className="mt-3"
                      style={{ display: "flex", flexWrap: "wrap", gap: "10px" }}
                    >
                      {newsPreviewUrls.map((url, idx) => (
                        <div
                          key={idx}
                          style={{
                            position: "relative",
                            display: "inline-block",
                          }}
                        >
                          <img
                            src={url}
                            alt={`news-preview-${idx}`}
                            style={{ width: "150px", borderRadius: "8px" }}
                          />
                          <button
                            type="button"
                            onClick={() => {
                              // ถ้า url นี้อยู่ใน existingNewsUrls ให้ลบออก
                              setExistingNewsUrls((prev) =>
                                prev.filter((u) => u !== url)
                              );
                              // ถ้าเป็นไฟล์ใหม่ ให้ลบออกจาก newsImages และ preview list
                              setNewsPreviewUrls((prev) =>
                                prev.filter((u) => u !== url)
                              );
                              // ลบไฟล์ที่ตรงกับ url นี้ (ถ้ามี)
                              setNewsImages((prevFiles) =>
                                prevFiles.filter((f) => URL.createObjectURL(f) !== url)
                              );
                            }}
                            style={{
                              position: "absolute",
                              top: "5px",
                              right: "5px",
                              background: "rgba(0,0,0,0.4)",
                              color: "white",
                              border: "none",
                              borderRadius: "50%",
                              width: "24px",
                              height: "24px",
                              fontSize: "18px",
                              cursor: "pointer",
                            }}
                            title="ลบภาพข่าว"
                          >
                            &times;
                          </button>
                        </div>
                      ))}
                    </div>
                  )}
                </div>

                <button
                  type="submit"
                  className="btn btn-primary"
                >
                  บันทึกการแก้ไข
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

// utils
function readFile(file) {
  return new Promise((resolve) => {
    const reader = new FileReader();
    reader.addEventListener("load", () => resolve(reader.result), false);
    reader.readAsDataURL(file);
  });
}

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
    image.addEventListener("error", (err) => reject(err));
    image.setAttribute("crossOrigin", "anonymous");
    image.src = url;
  });
}

export default Edit_News;
