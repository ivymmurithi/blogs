import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { HttpClient } from '@angular/common/http';

interface Blog {
  id: number;
  title: string;
  content: string;
  slug: string;
  created_at: string;
}

@Component({
  selector: 'app-blog-detail',
  templateUrl: './blog-detail.component.html',
})
export class BlogDetailComponent implements OnInit {
  blog?: Blog;

  constructor(private route: ActivatedRoute, private http: HttpClient) {}

  ngOnInit() {
    const slug = this.route.snapshot.paramMap.get('slug');
    this.http.get<Blog>(`http://localhost:8080/api/blog/${slug}`).subscribe(data => {
      this.blog = data;
    });
  }
}
