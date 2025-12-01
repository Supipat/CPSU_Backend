import React, { useEffect, useState } from "react";
import "../../css/admin/news.css";
import { Link } from "react-router-dom";
import axios from "axios";
import Headers from "../../component/header";
import Navbar from "../../component/navbar";
import Footer from "../../component/footer";
import Menu from "../../component/menu";

function News_admin() {
  const [news, setNews] = useState([]);

  useEffect(() => {
    const fetchNews = async () => {
      try {
        const res = await axios.get("http://localhost:8080/api/v1/admin/news");
        setNews(res.data);
      } catch (err) {
        console.error("Error fetching news", err);
      }
    };

    fetchNews();
  }, []);

  return (
    <>
      <Headers />
      <Navbar />

      <div className="container text-center mt-4">
        <div className="row">
          {/* Sidebar Menu */}
          <div className="col-sm-3">
            <Menu />
          </div>

          {/* Main Content */}
          <div className="col-sm-9">
            <div className="d-flex justify-content-between align-items-center mb-4">
              <h3 className="news-title">ข่าวสารทั้งหมด</h3>
              <Link to="/admin/addnews" className="btn-addnews">
                + เพิ่มข่าวสาร
              </Link>
            </div>

            {news.length === 0 ? (
              <p className="text-muted">ยังไม่มีข่าวสาร</p>
            ) : (
              <div className="row row-cols-1 row-cols-md-3 g-4">
                {news.map((item) => (
                  <div className="col" key={item.news_id}>
                    <Link
                      to={`/admin/news/${item.news_id}`}
                      className="text-decoration-none text-dark"
                    >
                      <div className="card h-100 shadow-sm news-card">
                        <img
                          src={
                            item?.cover_image
                              ? item.cover_image
                              : item?.images?.length > 0
                              ? item.images[0].file_image
                              : "/images/cpsu.png"
                          }
                          className="card-img-top news-img"
                          alt={item.title}
                        />
                        <div className="card-body">
                          <h5 className="card-title news-card-title">
                            {item.title.length > 70
                              ? item.title.slice(0, 70) + "..."
                              : item.title}
                          </h5>
                        </div>
                        <div className="card-footer">
                          <small className="text-muted">
                            {new Date(item.created_at).toLocaleString("th-TH", {
                              dateStyle: "short",
                              timeStyle: "short",
                            })}
                          </small>
                        </div>
                      </div>
                    </Link>
                  </div>
                ))}
              </div>
            )}
          </div>
        </div>
      </div>

      <div id="space"></div>
      <Footer />
    </>
  );
}

export default News_admin;
