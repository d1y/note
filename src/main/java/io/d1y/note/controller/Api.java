package io.d1y.note.controller;

import io.d1y.note.dao.NoteDao;
import io.d1y.note.repo.NoteRepo;
import io.d1y.note.utils.BaseRespone;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api")
public class Api {

  @Autowired
  private NoteRepo noteRepo;

  @GetMapping("/router/{routername}")
  public NoteDao catRouterContent(@PathVariable("routername") String router) {
    NoteDao data = noteRepo.findRouter(router);
    if (data == null) {
      return (new NoteDao());
    }
    return data;
  }

  @PostMapping("/router")
  @ResponseBody
  public BaseRespone update(@RequestParam("router") String router, @RequestParam("content") String content) {
    BaseRespone res = new BaseRespone();
    try {
      if (router.length() <= 0 || content.length() <= 0) {
        res.setCode(400);
        res.setMsg("参数传递错误");
      } else {
        // https://stackoverflow.com/a/61666895
        if (router.matches("[\\u4E00-\\u9FA5]+")) {
          res.setCode(404);
          res.setMsg("路由不应该为中文");
          return res;
        }
        NoteDao dao = noteRepo.findRouter(router);
        if (dao == null) {
          NoteDao noteData = new NoteDao();
          noteData.setRouter(router);
          noteData.setContent(content);
          noteData.setTitle("");
          noteRepo.save(noteData);
        } else {
          System.out.println("更新中: " + router);
          noteRepo.updateContent(content, router);
        }
        res.setCode(200);
        res.setMsg("更新成功");
      }
    } catch (Exception e) {
      res.setCode(500);
      res.setMsg(e.toString());
    }
    return res;
  }
}
