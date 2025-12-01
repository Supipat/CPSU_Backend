import React, { useEffect, useState } from "react";
import "../css/news_detail.css";
import { Link } from "react-router-dom";
import axios from "axios";
import Headers from "../component/header";
import Navbar from "../component/navbar";
import Footer from "../component/footer";

function News() {
  const [news, setNews] = useState([]);
  const [activeTab, setActiveTab] = useState("all");

  // กำหนดประเภทข่าว (static หรือจะดึงจาก API ก็ได้)
  const newsTypes = [
    { type_id: "all", type_name: "ข่าวทั้งหมด" },
    { type_id: 1, type_name: "ข่าวประชาสัมพันธ์" },
    { type_id: 2, type_name: "ทุนการศึกษา" },
    { type_id: 3, type_name: "รางวัลที่ได้รับ" },
    { type_id: 4, type_name: "กิจกรรมของภาควิชา" },
  ];

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

  // กรองข่าวตามแท็บที่เลือก
  const filteredNews =
    activeTab === "all"
      ? news
      : news.filter((item) => item.type_id === activeTab);

  function truncateText(text, maxLength = 35) {
    if (!text) return "";
    return text.length > maxLength ? text.slice(0, maxLength) + "..." : text;
  }

  return (
    <>
      <Headers />
      <Navbar />
      <div className="container text-center py-4">
        <div className="row mb-4">
          <div className="col">
            {/* Tabs */}
            <ul className="nav nav-tabs justify-content-center">
              {newsTypes.map((type) => (
                <li className="nav-item" key={type.type_id}>
                  <button
                    className={`nav-link ${
                      activeTab === type.type_id ? "active" : ""
                    }`}
                    onClick={() => setActiveTab(type.type_id)}
                  >
                    {type.type_name}
                  </button>
                </li>
              ))}
            </ul>
          </div>
        </div>

        <div className="row row-cols-1 row-cols-md-4 g-4">
          {filteredNews.length === 0 && (
            <p className="text-center">ไม่มีข่าวในหมวดนี้</p>
          )}
          {filteredNews.map((item, index) => (
            <div className="col" key={index}>
              <Link
                to={`/news/${item.news_id}`}
                className="text-decoration-none text-dark"
              >
                <div className="card h-100 ">
                  <img
                    src={
                      item.images && item.images.length > 0
                        ? item.images[0].file_image
                        : "/images/cpsu.png"
                    }
                    className="card-img-top"
                    alt={item.title}
                    style={{ height: "200px", objectFit: "cover" }}
                  />
                   <div className="card-body">
                        <h5 className="card-title" id="card-title">
                          {item.title}
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
      </div>
      <Footer />
    </>
  );
}

export default News;
