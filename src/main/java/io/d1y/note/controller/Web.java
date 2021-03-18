package io.d1y.note.controller;

import io.d1y.note.dao.NoteDao;
import io.d1y.note.repo.NoteRepo;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.util.Random;

@Controller
public class Web {

  @Autowired
  private NoteRepo noteRepo;

  /**
   * 渲染的路由的前缀
   */
  private final String webPrefix = "/views/";

  /**
   * 生成缓存随机最大值(最小长度为0)
   * @param max 最大长度
   * @return index
   */
  private int createCacheRandom(int max) {
    int min = 0;
    int randWord = new Random().nextInt(max - min + 1) + min;
    return randWord;
  }

  /**
   * 生成一个随机路由
   * TODO
   */
  private String createRandomPath() {

    String words = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
    int maxLen = createCacheRandom(10);
    String result = "";

    for (int i = 0; i < maxLen; i++) {
      int indexFly = createCacheRandom(words.length());
      result += words.charAt(indexFly);
    }

    return result;
  }

  @ResponseBody
  @RequestMapping("/")
  public void AsIndex(HttpServletResponse response) throws IOException {
    String rstring = createRandomPath();
    response.sendRedirect(webPrefix + rstring);
  }

  private final String routerName = "routerName";

  //
  // https://stackoverflow.com/questions/62820269/render-a-variable-as-html-in-java-spring-thymeleaf
  //
  @GetMapping(path = webPrefix + "{"+ routerName +"}")
  public String renderView(@PathVariable(routerName) String router, Model model) {
    boolean checkChinese = router.matches("[\\u4E00-\\u9FA5]+");
    boolean hasPush = false;
    String str = "";
    if (!checkChinese) {
      hasPush = true;
      NoteDao data = noteRepo.findRouter(router);
      if (data != null) {
        str = data.getContent();
      }
    } else {
      str = "路由错误";
    }
    model.addAttribute("hasPush", hasPush);
    model.addAttribute("content", str);
    model.addAttribute("title", router);
    return "index";
  }

}
