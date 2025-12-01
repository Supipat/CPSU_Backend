import React, { useEffect, useState } from "react";
import "../../css/admin/news_detail.css";
import Navbar from "../../component/navbar";
import Headers from "../../component/header";
import Footer from "../../component/footer";
import Menu from "../../component/menu";
import { useParams } from "react-router-dom";
import { Link } from "react-router-dom";
import axios from "axios";
import { useNavigate } from "react-router-dom";

const Detail_News = () => {
  const { id } = useParams();
  const [news, setNews] = useState(null);
  const navigate = useNavigate();

  const deleteNews = async () => {
    const confirmDelete = window.confirm("คุณแน่ใจหรือไม่ว่าต้องการลบข่าวนี้?");
    if (!confirmDelete) return;

    try {
      const response = await axios.delete(
        `http://localhost:8080/api/v1/admin/news/${id}`
      );
      console.log("News deleted successfully:", response.data);
      alert("ลบข่าวสารสำเร็จ");
      navigate("/admin/news");
    } catch (error) {
      console.error("Error deleting news:", error);
    }
  };
  useEffect(() => {
    const detailNews = async () => {
      try {
        const response = await axios.get(
          `http://localhost:8080/api/v1/admin/news/${id}`
        );
        const data = response.data;
        setNews(data);
      } catch (error) {
        console.log("Error fetching news details", error);
      }
    };
    detailNews();
  }, [id]);
  const images = news?.images || [];

  return (
    <>
      <Headers />
      <Navbar />
      <div className="container text-center">
        <div className="row">
          <div className="col-sm-3">
            <Menu />
          </div>
          <div className="col-sm-9">
            <div>
              <div id="group-btn-header-detail">
                <p id="news-name">{news?.title || "กำลังโหลด..."}</p>
                <div className="d-grid gap-2 d-md-flex justify-content-md-end">
                  <div className="edit_detele-position">
                    <Link
                      to={`/admin/editnews/${news?.news_id}`}
                      className="btn btn-warning"
                      id="btn-edit"
                    >
                      แก้ไข
                    </Link>
                    <div
                      className="btn btn-danger"
                      id="btn-delete"
                      onClick={deleteNews}
                    >
                      ลบ
                    </div>
                  </div>
                </div>
              </div>

              <div className="component_admin-detail-news row justify-content-between">
                <div className="col-2">
                  <label className="form-label" id="text-news-detail">
                    {news?.created_at
                      ? new Date(news.created_at).toLocaleString("th-TH", {
                          dateStyle: "short",
                          timeStyle: "short",
                        })
                      : ""}
                  </label>
                </div>
              </div>
              {images.length > 0 ? (
                <div
                  id="newsCarousel"
                  className="carousel slide mb-3"
                  data-bs-ride="carousel"
                >
                  <div className="carousel-inner">
                    {images.map((image, index) => (
                      <div
                        key={index}
                        className={`carousel-item ${
                          index === 0 ? "active" : ""
                        }`}
                      >
                        <img
                          src={image.file_image}
                          className="d-block w-100 img-news"
                          alt={`slide-${index}`}
                        />
                      </div>
                    ))}
                  </div>

                  {/* 👇 แสดงปุ่มควบคุมเฉพาะเมื่อมีมากกว่า 1 รูป */}
                  {images.length > 1 && (
                    <>
                      <button
                        className="carousel-control-prev"
                        type="button"
                        data-bs-target="#newsCarousel"
                        data-bs-slide="prev"
                      >
                        <span
                          className="carousel-control-prev-icon"
                          aria-hidden="true"
                        ></span>
                      </button>
                      <button
                        className="carousel-control-next"
                        type="button"
                        data-bs-target="#newsCarousel"
                        data-bs-slide="next"
                      >
                        <span
                          className="carousel-control-next-icon"
                          aria-hidden="true"
                        ></span>
                      </button>
                    </>
                  )}
                </div>
              ) : (
                <img
                  src="/images/cpsu.png"
                  className="img-news mb-3"
                  alt="default"
                />
              )}

              <div
                className="text-start mt-3"
                dangerouslySetInnerHTML={{ __html: news?.content || "" }}
              ></div>

              <div className="text-start mt-3">
                <p id="news-detail-more">รายละเอียดเพิ่มเติม</p>
                {news?.detail_url && news.detail_url !== "null" ? (
                  news.detail_url.split("\n").map((line, index) => (
                    <p key={index}>
                      <a href={line} target="_blank" rel="noopener noreferrer">
                        {line}
                      </a>
                    </p>
                  ))
                ) : (
                  <p>-</p>
                )}
              </div>
            </div>
          </div>
        </div>
      </div>
      <Footer />
    </>
  );
};

export default Detail_News;
